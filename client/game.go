package main

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
	"net/http"
	"splendor/utils"
	"sync"
)

type GameStatus struct {
	ConnectStatus string
	SessionID     string
	TableID       string
	PlayerID      string
	UserName      string
	HasJoin       bool
	*Client
	*GameCron
	ExitFunc func(interface{}) error
}

func ConstructGameStatus(c *Client) *GameStatus {
	g := &GameStatus{
		Client: c,
		GameCron: &GameCron{
			stop: make(chan struct{}),
			wg:   &sync.WaitGroup{},
		},
	}

	return g
}

func BuildErrorResponseError(code int) error {
	return errors.New(fmt.Sprintf("错误的返回码: %+v", code))
}

func (g *GameStatus) Info() []string {

	c := color.New(color.FgCyan)

	return []string{
		fmt.Sprintf("玩家[%+v] ", c.Sprint(g.UserName)),
		fmt.Sprintf("状态[%+v] ", c.Sprint(g.ConnectStatus)),
		fmt.Sprintf("会话[%+v] ", c.Sprint(utils.CompressUuid(g.SessionID))),
		fmt.Sprintf("房间[%+v] ", c.Sprint(utils.CompressUuid(g.TableID))),
		fmt.Sprintf("玩家[%+v] ", c.Sprint(utils.CompressUuid(g.PlayerID))),
	}
}

func (g *GameStatus) ShowPlayerInfo() {
	//black := color.New(color.BgBlack)
	fmt.Println("\033[40m                                                                                   \033[0m")
	infoRow := ""
	for _, info := range g.Info() {
		infoRow += info
	}
	fmt.Println(infoRow)
	fmt.Println("\033[40m                                                                                   \033[0m")
}

func (g *GameStatus) AskWhichTurn() (string, error) {
	resp, err := g.SendRequest("cur_player", map[string]any{})
	if err != nil {
		return "", err
	}

	content, err := ExtractBodyContent(resp)
	if err != nil {
		return "", err
	}

	return content, nil
}

func (g *GameStatus) JoinGame() (string, error) {
	resp, err := g.SendRequest("join", map[string]any{
		"username": "biang",
	})
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", BuildErrorResponseError(resp.StatusCode)
	}

	// 设置cookies
	g.Cookies = resp.Cookies()

	content, err := ExtractBodyContent(resp)
	if err != nil {
		return "", err
	}

	return content, nil
}

func (g *GameStatus) Alive() (string, error) {
	resp, err := g.SendRequest("alive", map[string]any{})
	if err != nil {
		return "", err
	}

	content, err := ExtractBodyContent(resp)
	if err != nil {
		return "", err
	}

	return content, nil
}

func (g *GameStatus) Leave() (string, error) {
	resp, err := g.SendRequest("leave", map[string]any{})
	if err != nil {
		return "", err
	}

	content, err := ExtractBodyContent(resp)
	if err != nil {
		return "", err
	}

	return content, nil
}

func (g *GameStatus) TableInfo() (string, error) {
	content, err := g.SendRequestAndGetContent("table_info", map[string]any{})
	if err != nil {
		return "", err
	}

	return content, nil
}

func (g *GameStatus) NextTurn() (string, error) {
	resp, err := g.SendRequest("next_turn", map[string]any{})
	if err != nil {
		return "", err
	}

	content, err := ExtractBodyContent(resp)
	if err != nil {
		return "", err
	}

	return content, nil
}

func (g *GameStatus) KeepALive() (string, error) {
	resp, err := g.SendRequest("keep_a_live", map[string]any{})
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("错误码: %+v", resp.StatusCode))
	}
	return "ok", nil
}

// IsOurTurn 检查是否为自己的回合，异常打印
func (g *GameStatus) IsOurTurn() bool {
	yes, err := g.IfMyTurn()
	if err != nil {
		fmt.Println(err)
		return false
	}
	if !yes {
		fmt.Println("不是你的回合")
		return false
	}
	return true
}

// IfMyTurn 判断是不是自己的回合
func (g *GameStatus) IfMyTurn() (bool, error) {
	content, err := g.AskWhichTurn()
	if err != nil {
		return false, err
	}
	fmt.Println(content)

	curPlayerID := gjson.Get(content, "current_player_id").String()
	if curPlayerID != g.PlayerID {
		return false, nil
	}
	return true, nil
}

// TakeThreeTokens 拿3个宝石
func (g *GameStatus) TakeThreeTokens(tokensString string) (string, error) {
	resp, err := g.SendRequest("take_three_tokens", map[string]any{
		"tokens": tokensString,
	})
	if err != nil {
		return "", err
	}
	content, err := ExtractBodyContent(resp)
	if err != nil {
		return "", err
	}
	ret := gjson.Get(content, "ret").Int()
	if ret != 0 {
		return "ret", nil
	}

	return "ok", nil
}

// TakeDoubleTokens 拿两个宝石
func (g *GameStatus) TakeDoubleTokens(tokenId int) (string, error) {
	resp, err := g.SendRequest("take_double_tokens", map[string]any{
		"token_id": tokenId,
	})
	if err != nil {
		return "", err
	}

	content, err := ExtractBodyContent(resp)
	if err != nil {
		return "", err
	}
	ret := gjson.Get(content, "ret").Int()
	if ret != 0 {
		return "ret", nil
	}

	return "ok", nil
}

// ReturnTokens 返还多余的宝石
func (g *GameStatus) ReturnTokens(tokensString string) (string, error) {

	resp, err := g.SendRequest("return_tokens", map[string]any{
		"tokens": tokensString,
	})
	if err != nil {
		return "", err
	}
	code, msg, err := CheckRespStatusCode(resp)
	if err != nil {
		return "", err
	}
	if code != http.StatusOK {
		return "", errors.New(msg)
	}

	return "ok", nil

}

var ExitError = errors.New("exit error")

func (g *GameStatus) CheckExit(signal interface{}) error {
	if g.ExitFunc == nil {
		return nil
	}
	return g.ExitFunc(signal)
}

func (g *GameStatus) SetExitFunc(f func(interface{}) error) *GameStatus {
	g.ExitFunc = f
	return g
}
