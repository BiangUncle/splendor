package model

import "time"

const RevealedDevelopmentCardPerLevel = 4

// Table 游戏桌面
type Table struct {
	Users                    []*Player              // 玩家
	GameTime                 time.Time              // 游戏时间
	DevelopmentCardStacks    *DevelopmentCardStacks // 发展卡堆
	RevealedDevelopmentCards *DevelopmentCardStacks // 暴露的发展卡
	NobleTilesStack          NobleTilesStack        // 贵族卡堆
	TokenStack               TokenStack             // 宝石卡堆

	CurrentUser *Player // 当前角色
}

// CreateANewTable 开一桌
func CreateANewTable() *Table {
	table := &Table{
		Users:                    nil,
		GameTime:                 time.Now(),
		DevelopmentCardStacks:    CreateANewDevelopmentCardStacks(),
		NobleTilesStack:          CreateANewNobleTilesStack(),
		TokenStack:               CreatANewTokenStack(),
		RevealedDevelopmentCards: CreateEmptyDevelopmentCardStacks(),
	}

	return table
}

func (t *Table) Shuffle() {
	t.DevelopmentCardStacks.Shuffle()
	t.NobleTilesStack.Shuffle()
}
