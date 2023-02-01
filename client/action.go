package main

import (
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"sort"
	"splendor/utils"
)

var ActionMap = map[int]func() error{
	1: DefaultAction.JoinGame,
}

var PreActionMap = map[int]func(status *GameStatus) error{
	1:  JoinGame,
	4:  TableInfo,
	5:  NextTurn,
	6:  TakeThreeTokens,
	7:  TakeDoubleTokens,
	8:  ReturnTokens,
	9:  PurchaseDevelopmentCard,
	10: ReserveDevelopmentCard,
	11: PurchaseHandCard,
	0:  Leave,
}

var PreActionName = map[int]Option{
	1:  Option{1, "加入游戏"},
	4:  Option{4, "桌面信息"},
	6:  Option{6, "三个宝石"},
	7:  Option{7, "两个宝石"},
	9:  Option{9, "购买发展卡"},
	10: Option{10, "保存发展卡"},
	11: Option{11, "购买手卡"},
	0:  Option{0, "离开"},
}

type Option struct {
	Index int
	Name  string
}

/*
MIND: 将所有的Action出现的报错打印集成在action中，不在外面进行打印
*/

type Action struct {
	gs       *GameStatus
	tryTime  int
	exitFunc func(interface{}) bool
}

var InvalidExitError = errors.New("手动退出循环错误")
var DefaultAction *Action

func InitDefaultAction() {
	DefaultAction = &Action{
		gs:       nil,
		tryTime:  0,
		exitFunc: nil,
	}
}

// ReturnTokens 返还宝石循环操作
func (a *Action) ReturnTokens() error {
	for {
		a.tryTime++
		tokensString := inputString("请输入要丢弃的宝石")

		if a.CheckExit(tokensString) {
			return InvalidExitError
		}
		_, err := a.gs.ReturnTokens(tokensString)
		if err != nil {
			fmt.Println(err)
			continue
		}
		return nil
	}
}

func (a *Action) SetExitFunc(f func(interface{}) bool) *Action {
	a.exitFunc = f
	return a
}

func (a *Action) CheckExit(signal interface{}) bool {
	if a.exitFunc == nil {
		return false
	}
	return a.exitFunc(signal)
}

func (a *Action) JoinGame() error {
	content, err := a.gs.JoinGame()
	if err != nil {
		fmt.Println(err)
		return err
	}
	a.gs.ConnectStatus = gjson.Get(content, "status").String()
	a.gs.TableID = gjson.Get(content, "table_id").String()
	a.gs.PlayerID = gjson.Get(content, "player_id").String()
	a.gs.SessionID = gjson.Get(content, "session_id").String()
	a.gs.UserName = gjson.Get(content, "username").String()

	go a.gs.RoutineKeepALive()

	return nil
}

func (a *Action) Alive() error {
	content, err := a.gs.Alive()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(content)
	return nil
}

func (a *Action) LoopAction(f func() error) {
	a.tryTime = 0
	for {
		a.tryTime++
		err := f()
		if err != nil {
			fmt.Println(err)
			continue
		}
		return
	}
}

func (a *Action) TestAction() error {
	i, err := utils.InputInt("")
	if err != nil {
		return err
	}
	if i == -1 {
		return errors.New("error")
	}
	fmt.Println("输入", i)
	return nil
}

func JoinGame(g *GameStatus) error {
	if g.HasJoin {
		fmt.Println("已经加入房间")
		return nil
	}

	content, err := g.JoinGame()
	if err != nil {
		return err
	}
	g.ConnectStatus = gjson.Get(content, "status").String()
	g.TableID = gjson.Get(content, "table_id").String()
	g.PlayerID = gjson.Get(content, "player_id").String()
	g.SessionID = gjson.Get(content, "session_id").String()
	g.UserName = gjson.Get(content, "username").String()

	go g.RoutineKeepALive()

	g.HasJoin = true
	return nil
}

func TableInfo(g *GameStatus) error {
	if !g.HasJoin {
		fmt.Println("还没加入房间")
		return nil
	}

	content, err := g.TableInfo()
	if err != nil {
		return err
	}
	tableInfo := gjson.Get(content, "tableInfo").String()
	fmt.Println(tableInfo)
	return nil
}

func Leave(g *GameStatus) error {
	if !g.HasJoin {
		fmt.Println("客户端: ", "退出游戏")
		return nil
	}
	content, err := g.Leave()
	if err != nil {
		return err
	}
	fmt.Println("客户端: ", "退出游戏")
	fmt.Println("服务端: ", content)
	g.Stop()
	g.wg.Wait()
	return nil
}

func NextTurn(g *GameStatus) error {
	if !g.IsOurTurn() {
		fmt.Println("当前不是你的回合")
		return nil
	}
	content, err := g.NextTurn()
	if err != nil {
		fmt.Println(err)
		return err
	}
	nextPlayerName := gjson.Get(content, "current_player_name").String()
	fmt.Println(fmt.Sprintf("下一个是 %+v 操作", nextPlayerName))
	return nil
}

func TakeThreeTokens(g *GameStatus) error {
	if !g.IsOurTurn() {
		fmt.Println("当前不是你的回合")
		return nil
	}
	tokensString, err := utils.InputString("")
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 确认是否强行退出
	if err = g.CheckExit(tokensString); err != nil {
		return err
	}
	msg, err := g.TakeThreeTokens(tokensString)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if msg == "ret" {
		LoopActionByFunction(ReturnTokens)(g)
	}
	return nil
}

func TakeDoubleTokens(g *GameStatus) error {
	if !g.IsOurTurn() {
		fmt.Println("当前不是你的回合")
		return nil
	}
	tokenId, err := utils.InputString("")
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 确认是否强行退出
	if err = g.CheckExit(tokenId); err != nil {
		return err
	}
	msg, err := g.TakeDoubleTokens(utils.ToInt(tokenId))
	if err != nil {
		fmt.Println(err)
		return err
	}
	if msg == "ret" {
		LoopActionByFunction(ReturnTokens)(g)
	}
	return nil
}

func ReturnTokens(g *GameStatus) error {
	tokensString := inputString("请输入要丢弃的宝石")

	if err := g.CheckExit(tokensString); err != nil {
		return err
	}

	_, err := g.ReturnTokens(tokensString)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func OptionsInfos() []string {
	var ret []string
	var options []Option
	for _, v := range PreActionName {
		options = append(options, v)
	}

	sort.Slice(options, func(i, j int) bool {
		return options[i].Index < options[j].Index
	})
	for _, op := range options {
		ret = append(ret, fmt.Sprintf("%d. [%s]", op.Index, op.Name))
	}

	return ret
}

func ShowOptionsInfos() string {
	options := OptionsInfos()
	ret := ""
	for _, option := range options {
		ret += option + ""
	}
	return ret
}

func LoopAction(action int) func(status *GameStatus) {
	a := &Action{}
	return func(g *GameStatus) {
		a.tryTime = 0
		for a.tryTime < 5 {
			a.tryTime++
			err := PreActionMap[action](g)
			if err == ExitError {
				fmt.Println("手动退出程序")
				return
			}
			if err != nil {
				fmt.Println(err)
				continue
			}
			return
		}
	}
}

func LoopActionByFunction(f func(status *GameStatus) error) func(status *GameStatus) {
	a := &Action{}
	return func(g *GameStatus) {
		a.tryTime = 0
		for a.tryTime < 5 {
			a.tryTime++
			err := f(g)
			if err == ExitError {
				fmt.Println("手动退出程序")
				return
			}
			if err != nil {
				fmt.Println(err)
				continue
			}
			return
		}
	}
}

func PurchaseDevelopmentCard(g *GameStatus) error {
	if !g.IsOurTurn() {
		fmt.Println("当前不是你的回合")
		return nil
	}
	tokensString, err := utils.InputString("请输入你支出的宝石")
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 确认是否强行退出
	if err = g.CheckExit(tokensString); err != nil {
		return err
	}
	cardIdx, err := utils.InputInt("请输入你购买的卡ID")
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 确认是否强行退出
	if cardIdx == -1 {
		return ExitError
	}

	_, err = g.PurchaseDevelopmentCard(cardIdx, tokensString)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func ReserveDevelopmentCard(g *GameStatus) error {
	if !g.IsOurTurn() {
		fmt.Println("当前不是你的回合")
		return nil
	}

	cardIdx, err := utils.InputInt("请输入你购买的卡ID")
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 确认是否强行退出
	if cardIdx == -1 {
		return ExitError
	}

	_, err = g.ReserveDevelopmentCard(cardIdx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func PurchaseHandCard(g *GameStatus) error {
	if !g.IsOurTurn() {
		fmt.Println("当前不是你的回合")
		return nil
	}

	tokensString, err := utils.InputString("请输入你支出的宝石")
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 确认是否强行退出
	if err = g.CheckExit(tokensString); err != nil {
		return err
	}
	cardIdx, err := utils.InputInt("请输入你购买的卡ID")
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 确认是否强行退出
	if cardIdx == -1 {
		return ExitError
	}

	_, err = g.PurchaseHandCard(cardIdx, tokensString)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
