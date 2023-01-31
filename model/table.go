package model

import (
	"errors"
	"fmt"
	"splendor/utils"
	"time"
)

const RevealedDevelopmentCardNumPerLevel = 4

var GlobalTable = make(map[string]*Table)
var defaultTable *Table
var defaultTableID string

// Table 游戏桌面
type Table struct {
	Players  []*Player // 玩家
	GameTime time.Time // 游戏时间

	DevelopmentCardStacks    *DevelopmentCardStacks // 发展卡堆
	RevealedDevelopmentCards *DevelopmentCardStacks // 暴露的发展卡

	NobleTilesStack    NobleTilesStack // 贵族卡堆
	RevealedNobleTiles NobleTilesStack // 暴露的贵族卡

	TokenStack TokenStack // 宝石卡堆

	CurrentPlayer    *Player // 当前角色
	CurrentPlayerIdx int     // 当前角色索引

	TableID string // 桌台ID
	Name    string // 桌台名字
}

// InitDefaultTable 创建一个default桌台
func InitDefaultTable() {
	table, tableID, err := JoinNewTable()
	if err != nil {
		panic(err)
	}
	table.Name = "默认桌台"

	defaultTable = table
	defaultTableID = tableID

	defaultTable.Reveal()
}

// JoinDefaultTable 加入default桌台
func JoinDefaultTable(player *Player) (*Table, string, error) {
	nextIdx := len(defaultTable.Players)
	player.Name = fmt.Sprintf("玩家[%+v]", nextIdx)
	defaultTable.AddPlayer(player)
	return defaultTable, defaultTableID, nil
}

// CreateTable 创建一个桌布对象
func CreateTable() *Table {
	table := &Table{
		Players:                  make([]*Player, 0),
		GameTime:                 time.Now(),
		DevelopmentCardStacks:    CreateANewDevelopmentCardStacks(),
		NobleTilesStack:          CreateANewNobleTilesStack(),
		TokenStack:               CreatANewTokenStack(),
		RevealedDevelopmentCards: CreateEmptyDevelopmentCardStacks(),
		RevealedNobleTiles:       make([]*NobleTile, 3),
	}

	return table
}

// CreateTableID 创建桌台ID
func CreateTableID() string {
	return utils.GetUuidV4()
}

// JoinNewTable 加入一个新的桌台
func JoinNewTable() (*Table, string, error) {

	table := CreateTable()
	tableID := CreateTableID()
	table.TableID = tableID

	if _, ok := GlobalTable[tableID]; ok {
		return nil, "", errors.New(fmt.Sprintf("这个桌子已经有了"))
	}
	GlobalTable[tableID] = table
	return table, tableID, nil
}

// GetGlobalTable 根据tableID获取桌台
func GetGlobalTable(tableID string) (*Table, error) {
	if table, ok := GlobalTable[tableID]; ok {
		return table, nil
	}
	return nil, errors.New(fmt.Sprintf("没有这个ID的桌, tableID = %+v", tableID))
}

func GetDefaultTable() *Table {
	return defaultTable
}

// LeaveDefaultTable 离开default桌子
func LeaveDefaultTable(playerId string) {
	leaveId := -1
	var leavePlayer *Player
	for idx, player := range defaultTable.Players {
		if player.PlayerID == playerId {
			leaveId = idx
			leavePlayer = player
			break
		}
	}

	// 该桌台里面有需要退出的玩家
	if leaveId != -1 {
		defaultTable.Players = append(defaultTable.Players[:leaveId], defaultTable.Players[leaveId+1:]...)
		// 如果为空，全部置空
		if len(defaultTable.Players) == 0 {
			defaultTable.CurrentPlayerIdx = -1
			defaultTable.CurrentPlayer = nil
		} else if defaultTable.CurrentPlayer == leavePlayer { // 如果当前操作对象掉线, 下一个为操作对象
			n := len(defaultTable.Players)
			defaultTable.CurrentPlayerIdx = (defaultTable.CurrentPlayerIdx + 1) % n
		}
	}
}

// ClearDefaultTable 清理default桌子
func ClearDefaultTable() {
	defaultTable.ClearLoginOutPlayer()
}

// AddPlayer 添加玩家
func (t *Table) AddPlayer(p *Player) {
	t.Players = append(t.Players, p)
	if len(t.Players) == 1 {
		t.CurrentPlayer = p
		t.CurrentPlayerIdx = 0
	}
}

// Shuffle 发牌
func (t *Table) Shuffle() {
	t.DevelopmentCardStacks.Shuffle()
	t.NobleTilesStack.Shuffle()
}

// RevealDevelopmentCard 初始化发展卡
func (t *Table) RevealDevelopmentCard() (err error) {
	topCards, err := t.DevelopmentCardStacks.TopStack.TakeTopNCard(RevealedDevelopmentCardNumPerLevel)
	if err != nil {
		return
	}
	t.RevealedDevelopmentCards.TopStack = topCards

	middleCards, err := t.DevelopmentCardStacks.MiddleStack.TakeTopNCard(RevealedDevelopmentCardNumPerLevel)
	if err != nil {
		return
	}
	t.RevealedDevelopmentCards.MiddleStack = middleCards

	bottomCards, err := t.DevelopmentCardStacks.BottomStack.TakeTopNCard(RevealedDevelopmentCardNumPerLevel)
	if err != nil {
		return
	}
	t.RevealedDevelopmentCards.BottomStack = bottomCards

	return nil
}

// RevealNobleTiles 初始化贵族卡
func (t *Table) RevealNobleTiles() (err error) {

	userNum := len(t.Players)

	nobleTiles, err := t.NobleTilesStack.TakeTopNCard(userNum + 1)
	if err != nil {
		return
	}
	t.RevealedNobleTiles = nobleTiles

	return nil
}

// Reveal 展示发展卡和贵族卡
func (t *Table) Reveal() (err error) {
	err = t.RevealDevelopmentCard()
	if err != nil {
		return
	}
	err = t.RevealNobleTiles()
	if err != nil {
		return
	}
	return nil
}

// IsExistRevealedDevelopmentCard 判断这个卡是否在场上
func (t *Table) IsExistRevealedDevelopmentCard(idx int) bool {
	return t.RevealedDevelopmentCards.IsExistCard(idx)
}

// ReplaceRevealedDevelopmentCard 补充场上的发展卡
// todo: 可优化，有大量重复逻辑
func (t *Table) ReplaceRevealedDevelopmentCard(cardLevel int) error {

	switch cardLevel {
	case DevelopmentCardLevelTop:
		card, err := t.DevelopmentCardStacks.TopStack.TakeTopCard()
		// 如果没卡就不补充
		if err != nil && len(t.DevelopmentCardStacks.TopStack) == 0 {
			return nil
		}
		if err != nil {
			return err
		}
		err = t.RevealedDevelopmentCards.TopStack.PutNewCardToEmptySite(card)
		if err != nil {
			return err
		}
		return nil
	case DevelopmentCardLevelMiddle:
		card, err := t.DevelopmentCardStacks.MiddleStack.TakeTopCard()
		// 如果没卡就不补充
		if err != nil && len(t.DevelopmentCardStacks.MiddleStack) == 0 {
			return nil
		}
		if err != nil {
			return err
		}
		err = t.RevealedDevelopmentCards.MiddleStack.PutNewCardToEmptySite(card)
		if err != nil {
			return err
		}
		return nil
	case DevelopmentCardLevelBottom:
		card, err := t.DevelopmentCardStacks.BottomStack.TakeTopCard()
		// 如果没卡就不补充
		if err != nil && len(t.DevelopmentCardStacks.BottomStack) == 0 {
			return nil
		}
		if err != nil {
			return err
		}
		err = t.RevealedDevelopmentCards.BottomStack.PutNewCardToEmptySite(card)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New(fmt.Sprintf("你这个发展卡等级好像有问题，等级 = %d", cardLevel))
}

// RemoveRevealedNoble 移除贵族
func (t *Table) RemoveRevealedNoble(idx int) error {
	if len(t.RevealedNobleTiles) <= idx {
		return errors.New(fmt.Sprintf("好像没有这么多贵族， 目前长度 %d, 你要移除第 %d 个", len(t.RevealedNobleTiles), idx))
	}
	t.RevealedNobleTiles[idx] = nil
	return nil
}

// TableInfoString 玩家的信息
func (t *Table) TableInfoString() []string {
	return []string{
		fmt.Sprintf("%-15s %+v", "Name:", t.Name),
		fmt.Sprintf("%-15s %+v", "CurPlayer:", t.CurrentPlayer.Name),
		fmt.Sprintf("%-15s %+v", "Token:", t.TokenStack),
		fmt.Sprintf("%-15s %+v", "DevCard:", t.DevelopmentCardStacks.ShowIdxInfo()),
		fmt.Sprintf("%-15s %+v", "RevealCards:", t.RevealedDevelopmentCards.ShowIdxInfo()),
		fmt.Sprintf("%-15s %+v", "Noble:", t.NobleTilesStack.ShowIdxInfo()),
		fmt.Sprintf("%-15s %+v", "RevealNoble:", t.RevealedNobleTiles.ShowIdxInfo()),
	}
}

// ShowInfo 展示信息
func (t *Table) ShowInfo() {
	fmt.Printf("|=========Table=========\n")
	fmt.Printf("| Token: %+v\n", t.TokenStack)
	fmt.Printf("| DevCard: %+v\n", t.DevelopmentCardStacks.ShowIdxInfo())
	fmt.Printf("| RevealCards: %+v\n", t.RevealedDevelopmentCards.ShowIdxInfo())
	fmt.Printf("| Noble: %+v\n", t.NobleTilesStack.ShowIdxInfo())
	fmt.Printf("| RevealNoble: %+v\n", t.RevealedNobleTiles.ShowIdxInfo())
	fmt.Printf("|=======================\n")
}

// ShowTableInfo 展示整场游戏信息
func (t *Table) ShowTableInfo() string {

	infos := make([][]string, 0)

	ret := ""

	for _, player := range t.Players {
		infos = append(infos, player.PlayerInfoString())
	}

	line := ""
	for j := 0; j < len(infos); j++ {
		line = line + fmt.Sprintf("%s", "\u001B[40m                                \u001B[0m")
	}

	left := "\u001B[40m \u001B[0m"

	ret = ret + fmt.Sprintf("%s\n", line)

	for i := 0; i < len(infos[0]); i++ {
		infoRow := ""
		for j := 0; j < len(infos); j++ {
			infoRow = infoRow + fmt.Sprintf("%s %-30s", left, infos[j][i])
		}
		infoRow = infoRow + "\n"
		ret = ret + infoRow
	}

	ret = ret + fmt.Sprintf("%s\n", line)

	for _, info := range t.TableInfoString() {
		ret = ret + fmt.Sprintf("%s %s\n", left, info)
	}

	ret = ret + fmt.Sprintf("%s\n", line)

	return ret
}

// NextTurn 下一位
func (t *Table) NextTurn() *Player {
	n := len(t.Players)
	t.CurrentPlayerIdx = (t.CurrentPlayerIdx + 1) % n
	t.CurrentPlayer = t.Players[t.CurrentPlayerIdx]
	return t.CurrentPlayer
}

func (t *Table) FindPlayerIdx(p *Player) int {
	ret := -1
	for idx, player := range t.Players {
		if player == p {
			ret = idx
		}
	}
	return ret
}

// ClearLoginOutPlayer 清楚未在线用户
func (t *Table) ClearLoginOutPlayer() {
	var onlinePlayers []*Player
	for _, player := range t.Players {
		if CheckIfOnline(player.PlayerID) {
			onlinePlayers = append(onlinePlayers, player)
		} else {
			fmt.Printf("玩家: %+v 已离线\n", player.Name)
		}
	}

	beforeCurrentPlayer := t.CurrentPlayer

	t.CurrentPlayer = nil
	t.CurrentPlayerIdx = -1
	t.Players = onlinePlayers
	for idx, player := range t.Players {
		if player == beforeCurrentPlayer {
			t.CurrentPlayerIdx = idx
		}
	}

	// todo 有BUG，有人最好直接掉线
	if len(t.Players) > 0 {
		t.CurrentPlayerIdx = 0
		t.CurrentPlayer = t.Players[0]
	}
}

func (t *Table) ShowVisualInfo() {

	fmt.Println("\033[40m                                                                \033[0m")
	fmt.Println(t.TokenStack.Visual())
	fmt.Printf("[%d] %s\n", len(t.DevelopmentCardStacks.TopStack), t.RevealedDevelopmentCards.TopStack.Visual())
	fmt.Printf("[%d] %s\n", len(t.DevelopmentCardStacks.MiddleStack), t.RevealedDevelopmentCards.MiddleStack.Visual())
	fmt.Printf("[%d] %s\n", len(t.DevelopmentCardStacks.BottomStack), t.RevealedDevelopmentCards.BottomStack.Visual())
	fmt.Println(t.RevealedNobleTiles.Visual())
	fmt.Println("\033[40m                                                                \033[0m")
}
