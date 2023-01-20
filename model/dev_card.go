package model

import (
	"fmt"
	"math/rand"
)

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

type DevelopmentCardStack []*DevelopmentCard

// DevelopmentCard 发展卡
type DevelopmentCard struct {
	Idx       int        // 唯一索引
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

func CreateEmptyDevelopmentCardStacks() *DevelopmentCardStacks {
	return &DevelopmentCardStacks{
		TopStack:    make(DevelopmentCardStack, RevealedDevelopmentCardPerLevel),
		MiddleStack: make(DevelopmentCardStack, RevealedDevelopmentCardPerLevel),
		BottomStack: make(DevelopmentCardStack, RevealedDevelopmentCardPerLevel),
	}
}

func (s DevelopmentCardStack) Copy() DevelopmentCardStack {
	cpy := make(DevelopmentCardStack, len(s))
	copy(cpy, s)
	return cpy
}

// Copy 复制
func (s *DevelopmentCardStacks) Copy() *DevelopmentCardStacks {

	cpy := &DevelopmentCardStacks{
		TopStack:    s.TopStack.Copy(),
		MiddleStack: s.MiddleStack.Copy(),
		BottomStack: s.BottomStack.Copy(),
	}

	return cpy
}

// Shuffle 打乱牌堆
func (s DevelopmentCardStack) Shuffle() {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

func (s DevelopmentCardStack) ShowIdxInfo() {
	idxInfo := make([]int, len(s))
	for i, card := range s {
		idxInfo[i] = card.Idx
	}
	fmt.Printf("%+v\n", idxInfo)
}

// Shuffle 打乱牌堆
func (s *DevelopmentCardStacks) Shuffle() {
	s.TopStack.Shuffle()
	s.MiddleStack.Shuffle()
	s.BottomStack.Shuffle()
}

// ShowIdxInfo 打乱牌堆
func (s *DevelopmentCardStacks) ShowIdxInfo() {
	s.TopStack.ShowIdxInfo()
	s.MiddleStack.ShowIdxInfo()
	s.BottomStack.ShowIdxInfo()
}