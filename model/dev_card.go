package model

import "math/rand"

const DevelopmentCardNumber = 90     // 总共的卡牌数量
const DevelopmentCardLevelNumber = 3 // 发展卡等级数量

const (
	DevelopmentCardLevelTop    = iota + 1 // 一级
	DevelopmentCardLevelMiddle            // 二级
	DevelopmentCardLevelBottom            // 三级
)

const (
	TopLevelNumber    = 20 // 三级卡数量
	MiddleLevelNumber = 30 // 二级卡数量
	BottomLevelNumber = 40 // 一级卡数量
)

var defaultDevelopmentCardStacks *DevelopmentCardStacks

type DevelopmentCardStack []DevelopmentCard

// DevelopmentCard 发展卡
type DevelopmentCard struct {
	Level     int        // 等级
	BonusType int        // 奖励类别
	Prestige  int        // 声望
	Acquires  TokenStack // 所需的宝石列表
}

// DevelopmentCardStacks 发展卡堆
type DevelopmentCardStacks struct {
	TopStack    DevelopmentCardStack // 三级卡堆
	MiddleStack DevelopmentCardStack // 二级卡堆
	BottomStack DevelopmentCardStack // 一级卡堆
}

// CreateANewDevelopmentCardStacks 创建一个新的贵族堆
func CreateANewDevelopmentCardStacks() *DevelopmentCardStacks {
	return defaultDevelopmentCardStacks.Copy()
}

func (s DevelopmentCardStack) Copy() DevelopmentCardStack {
	cpy := make(DevelopmentCardStack, len(s))
	copy(cpy, s)
	return cpy
}

// Copy 复制
func (d *DevelopmentCardStacks) Copy() *DevelopmentCardStacks {

	cpy := &DevelopmentCardStacks{
		TopStack:    d.TopStack.Copy(),
		MiddleStack: d.MiddleStack.Copy(),
		BottomStack: d.BottomStack.Copy(),
	}

	return cpy
}

func ShuffleDevelopmentCard(developmentCards []DevelopmentCard) []DevelopmentCard {
	rand.Shuffle(len(developmentCards), func(i, j int) {
		developmentCards[i], developmentCards[j] = developmentCards[j], developmentCards[i]
	})
	return developmentCards
}
