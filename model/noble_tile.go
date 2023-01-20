package model

import (
	"errors"
	"fmt"
	"math/rand"
)

const NobleTitleNumber = 10

type NobleTile struct {
	Idx      int        // 唯一索引
	Prestige int        // 声望
	Acquires TokenStack // 所需的宝石列表
}

var defaultNobleTilesStack NobleTilesStack

type NobleTilesStack []*NobleTile

// CreateANewNobleTilesStack 创建一个新的贵族堆
func CreateANewNobleTilesStack() NobleTilesStack {
	return defaultNobleTilesStack.Copy()
}

// Copy 复制
func (s NobleTilesStack) Copy() NobleTilesStack {
	cpy := make(NobleTilesStack, len(s))
	copy(cpy, s)
	return cpy
}

// Shuffle 打乱牌堆
func (s NobleTilesStack) Shuffle() {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

// ShowIdxInfo 展示信息
func (s NobleTilesStack) ShowIdxInfo() string {
	idxInfo := make([]int, len(s))
	for i, noble := range s {
		idxInfo[i] = noble.Idx
	}
	//fmt.Printf("%+v\n", idxInfo)
	return fmt.Sprintf("%+v", idxInfo)
}

// TakeTopCard 翻第一张牌
func (s *NobleTilesStack) TakeTopCard() (*NobleTile, error) {

	ret, err := s.TakeTopNCard(1)
	if err != nil {
		return nil, err
	}
	return ret[0], nil

}

// TakeTopNCard 翻顶上 n 张牌
func (s *NobleTilesStack) TakeTopNCard(n int) (NobleTilesStack, error) {
	if len(*s) < n {
		return nil, errors.New(fmt.Sprintf("没有牌可以翻， 需要 %d 张，但是只有 %d 张。", n, len(*s)))
	}

	ret := (*s)[:n]
	*s = (*s)[n+1:]
	return ret, nil
}
