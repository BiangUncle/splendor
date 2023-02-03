package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"splendor/utils"
	"sync"
	"time"
)

const RevealedDevelopmentCardNumPerLevel = 4
const WinPrestigeCondition = 15

var GlobalTable = make(map[string]*Table)
var defaultTable *Table
var defaultTableID string

// Table 游戏桌面
type Table struct {
	Players                  []*Player              `json:"players"`                    // 玩家
	PlayersLock              *sync.Mutex            `json:"players_lock"`               // 玩家操作锁
	GameTime                 time.Time              `json:"game_time"`                  // 游戏时间
	DevelopmentCardStacks    *DevelopmentCardStacks `json:"development_card_stacks"`    // 发展卡堆
	RevealedDevelopmentCards *DevelopmentCardStacks `json:"revealed_development_cards"` // 暴露的发展卡
	NobleTilesStack          NobleTilesStack        `json:"noble_tiles_stack"`          // 贵族卡堆
	RevealedNobleTiles       NobleTilesStack        `json:"revealed_noble_tiles"`       // 暴露的贵族卡
	TokenStack               TokenStack             `json:"token_stack"`                // 宝石卡堆
	CurrentPlayer            *Player                `json:"current_player"`             // 当前角色
	CurrentPlayerIdx         int                    `json:"current_player_idx"`         // 当前角色索引
	TableID                  string                 `json:"table_id"`                   // 桌台ID
	Name                     string                 `json:"name"`                       // 桌台名字
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

	utils.SystemPrintf("玩家: %s 加入房间: %s\n", player.Name, defaultTable.Name)

	return defaultTable, defaultTableID, nil
}

// CreateTable 创建一个桌布对象
func CreateTable() *Table {
	table := &Table{
		Players:                  make([]*Player, 0),
		PlayersLock:              &sync.Mutex{},
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

	fmt.Printf("[SYSTEM] 玩家: %s 退出房间: %s\n", leavePlayer.Name, defaultTable.Name)
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

// NextTurn 下一位
func (t *Table) NextTurn() *Player {
	n := len(t.Players)

	t.PlayersLock.Lock()
	defer t.PlayersLock.Unlock()

	t.CurrentPlayer.ShowPlayerInfo()
	t.CurrentPlayerIdx = (t.CurrentPlayerIdx + 1) % n
	t.CurrentPlayer = t.Players[t.CurrentPlayerIdx]
	t.CurrentPlayer.ShowPlayerInfo()

	return t.CurrentPlayer
}

// FindPlayerIdx 确认角色在桌面的索引
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

	var offlinePlayers []*Player
	for _, player := range t.Players {
		if !CheckIfOnline(player.PlayerID) {
			offlinePlayers = append(offlinePlayers, player)
			fmt.Printf("玩家: %+v 已离线\n", player.Name)
		}
	}

	if len(offlinePlayers) == 0 {
		fmt.Println("没人掉线")
		return
	}

	for _, player := range offlinePlayers {
		t.RemovePlayer(player)
	}
}

// RemovePlayerByIdx 通过索引删除角色
func (t *Table) RemovePlayerByIdx(idx int) {
	if idx < 0 || idx >= len(t.Players) {
		return
	}
	player := t.Players[idx]
	t.RemovePlayer(player)
}

// RemovePlayer 桌面移除玩家
func (t *Table) RemovePlayer(p *Player) {
	if len(t.Players) == 0 {
		return
	}
	t.PlayersLock.Lock()
	defer t.PlayersLock.Unlock()

	curPlayer := t.CurrentPlayer
	if p == curPlayer { // 如果当前玩家为需要删除的玩家，则选择下一个玩家作为当前玩家
		curPlayer = t.Players[t.NextPlayerIdx()]
	}

	// 创建后续的玩家
	var newPlayers []*Player
	for _, player := range t.Players {
		if player != p {
			newPlayers = append(newPlayers, player)
		}
	}

	t.Players = newPlayers

	for idx, player := range t.Players {
		if player == curPlayer {
			t.CurrentPlayer = player
			t.CurrentPlayerIdx = idx
			break
		}
	}

	return
}

// NextPlayerIdx 下一个执行玩家的索引
func (t *Table) NextPlayerIdx() int {
	if len(t.Players) == 0 {
		return -1
	}
	n := len(t.Players)
	nextPlayerIdx := (t.CurrentPlayerIdx + 1) % n
	return nextPlayerIdx
}

type TableStatus struct {
	TokenStatus       TokenStack `json:"token_status"`
	TopStackStatus    []int      `json:"top_stack_status"`
	MiddleStackStatus []int      `json:"middle_stack_status"`
	BottomStackStatus []int      `json:"bottom_stack_status"`
	NobleTileStatus   []int      `json:"noble_tile_status"`
}

func (t *Table) Status() (string, error) {
	s := &TableStatus{}
	s.TopStackStatus = t.RevealedDevelopmentCards.TopStack.Status()
	s.MiddleStackStatus = t.RevealedDevelopmentCards.MiddleStack.Status()
	s.BottomStackStatus = t.RevealedDevelopmentCards.BottomStack.Status()
	s.NobleTileStatus = t.RevealedNobleTiles.Status()
	s.TokenStatus = t.TokenStack

	b, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (t *Table) LoadStatus(s string) error {
	status := &TableStatus{}
	err := json.Unmarshal([]byte(s), status)
	if err != nil {
		return err
	}
	err = t.RevealedDevelopmentCards.TopStack.LoadStatus(status.TopStackStatus)
	if err != nil {
		return err
	}
	err = t.RevealedDevelopmentCards.MiddleStack.LoadStatus(status.MiddleStackStatus)
	if err != nil {
		return err
	}
	err = t.RevealedDevelopmentCards.BottomStack.LoadStatus(status.BottomStackStatus)
	if err != nil {
		return err
	}
	err = t.RevealedNobleTiles.LoadStatus(status.NobleTileStatus)
	if err != nil {
		return err
	}
	t.TokenStack = status.TokenStatus

	return nil
}

func (t *Table) WhoWin() (*Player, int) {
	var winPlayer *Player
	var winIdx int
	for idx, player := range t.Players {
		if player.Prestige > WinPrestigeCondition {
			if winPlayer != nil && winPlayer.Prestige > player.Prestige {
				continue
			}
			if winPlayer != nil &&
				winPlayer.Prestige == player.Prestige &&
				winPlayer.CardNum() <= player.CardNum() {
				continue
			}
			winPlayer = player
			winIdx = idx
		}
	}

	return winPlayer, winIdx
}
