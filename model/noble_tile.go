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
var NobleTilesMap = make(map[int]*NobleTile, 0)

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

func (s NobleTilesStack) Status() []int {
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

func (s *NobleTilesStack) LoadStatus(status []int) error {
	*s = make(NobleTilesStack, len(status))
	for i, idx := range status {
		if _, ok := NobleTilesMap[idx]; ok {
			(*s)[i] = NobleTilesMap[idx]
		} else {
			(*s)[i] = nil
		}
	}
	return nil
}
