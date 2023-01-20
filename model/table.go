package model

import "time"

// Table 游戏桌面
type Table struct {
	Users                []*Player              // 玩家
	GameTime             time.Time              // 游戏时间
	DevelopmentCardStack *DevelopmentCardStacks // 发展卡堆
	NobleTilesStack      NobleTilesStack        // 贵族卡堆
	TokenStack           TokenStack             // 宝石卡堆

	CurrentUser *Player // 当前角色
}

// CreateANewTable 开一桌
func CreateANewTable() *Table {
	table := &Table{
		Users:                nil,
		GameTime:             time.Now(),
		DevelopmentCardStack: CreateANewDevelopmentCardStacks(),
		NobleTilesStack:      CreateANewNobleTilesStack(),
		TokenStack:           CreatANewTokenStack(),
	}

	return table
}
