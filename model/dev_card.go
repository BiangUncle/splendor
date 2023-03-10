package model

import (
	"errors"
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
var DevelopmentCardMap = make(map[int]*DevelopmentCard, 0)

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
		TopStack:    make(DevelopmentCardStack, RevealedDevelopmentCardNumPerLevel),
		MiddleStack: make(DevelopmentCardStack, RevealedDevelopmentCardNumPerLevel),
		BottomStack: make(DevelopmentCardStack, RevealedDevelopmentCardNumPerLevel),
	}
}

func DevelopmentCardIndexTransfer(cardIdx int) int {
	if cardIdx < 10000 {
		cardIdx += 10000
	}
	return cardIdx
}

// ToTokenStack 发展卡列表转成token列表
func (s DevelopmentCardStack) ToTokenStack() TokenStack {
	tokens := make(TokenStack, TokenTypeNumber)

	for _, card := range s {
		tokens[card.BonusType]++
	}

	return tokens
}

func (s DevelopmentCardStack) Copy() DevelopmentCardStack {
	cpy := make(DevelopmentCardStack, len(s))
	copy(cpy, s)
	return cpy
}

// Shuffle 打乱牌堆
func (s DevelopmentCardStack) Shuffle() {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

// IsExistCard 牌堆中是否有这张牌
func (s DevelopmentCardStack) IsExistCard(cardIdx int) bool {
	for _, card := range s {
		if card == nil {
			continue
		}
		if card.Idx == cardIdx {
			return true
		}
	}
	return false
}

// TakeCard 拿走一张卡牌，并置空
func (s *DevelopmentCardStack) TakeCard(cardIdx int) (*DevelopmentCard, bool) {
	selectedIdx := -1

	for idx, card := range *s {
		if card == nil {
			continue
		}
		if card.Idx == cardIdx {
			selectedIdx = idx
			break
		}
	}

	if selectedIdx != -1 {
		ret := (*s)[selectedIdx]
		(*s)[selectedIdx] = nil
		return ret, true
	}

	return nil, false
}

// PutNewCardToEmptySite 把一张新的卡放在空的位置
func (s *DevelopmentCardStack) PutNewCardToEmptySite(newCard *DevelopmentCard) error {

	selectedIdx := -1

	for idx, card := range *s {
		if card == nil {
			selectedIdx = idx
			break
		}
	}

	if selectedIdx == -1 {
		return errors.New(fmt.Sprintf("这个地方没有地方可以放卡呀，状态：%+v。", *s))
	}

	(*s)[selectedIdx] = newCard

	return nil
}

// TakeTopCard 翻第一张牌
func (s *DevelopmentCardStack) TakeTopCard() (*DevelopmentCard, error) {

	ret, err := s.TakeTopNCard(1)
	if err != nil {
		return nil, err
	}
	return ret[0], nil

}

// TakeTopNCard 翻顶上 n 张牌
func (s *DevelopmentCardStack) TakeTopNCard(n int) (DevelopmentCardStack, error) {
	if len(*s) < n {
		return nil, errors.New(fmt.Sprintf("没有牌可以翻， 需要 %d 张，但是只有 %d 张。", n, len(*s)))
	}

	ret := (*s)[:n]
	*s = (*s)[n:]
	return ret, nil
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
func (s *DevelopmentCardStacks) Shuffle() {
	s.TopStack.Shuffle()
	s.MiddleStack.Shuffle()
	s.BottomStack.Shuffle()
}

// IsExistCard 牌堆中是否有这张牌
func (s *DevelopmentCardStacks) IsExistCard(cardIdx int) bool {
	return s.TopStack.IsExistCard(cardIdx) || s.MiddleStack.IsExistCard(cardIdx) || s.BottomStack.IsExistCard(cardIdx)
}

// TakeCard 拿走一张卡牌，并置空
func (s *DevelopmentCardStacks) TakeCard(cardIdx int) (*DevelopmentCard, int, bool) {
	card, success := s.TopStack.TakeCard(cardIdx)
	if success {
		return card, DevelopmentCardLevelTop, true
	}
	card, success = s.MiddleStack.TakeCard(cardIdx)
	if success {
		return card, DevelopmentCardLevelMiddle, true
	}
	card, success = s.BottomStack.TakeCard(cardIdx)
	if success {
		return card, DevelopmentCardLevelBottom, true
	}
	return nil, -1, false
}

func (s DevelopmentCardStack) Status() []int {
	var ret []int
	for _, c := range s {
		if c == nil {
			ret = append(ret, 0)
		} else {
			ret = append(ret, c.Idx)
		}
	}
	return ret
}

func (s *DevelopmentCardStack) LoadStatus(status []int) error {
	*s = make(DevelopmentCardStack, RevealedDevelopmentCardNumPerLevel)
	for i, idx := range status {
		if _, ok := DevelopmentCardMap[idx]; ok {
			(*s)[i] = DevelopmentCardMap[idx]
		} else {
			(*s)[i] = nil
		}
	}
	return nil
}

func (s DevelopmentCardStacks) Status() [][]int {
	var ret [][]int
	ret = append(ret, s.TopStack.Status())
	ret = append(ret, s.MiddleStack.Status())
	ret = append(ret, s.BottomStack.Status())
	return ret
}
