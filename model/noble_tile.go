package model

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
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
		if noble == nil {
			idxInfo[i] = -1
		} else {
			idxInfo[i] = noble.Idx % 100
		}
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

func (n *NobleTile) Visual() string {
	require := ""
	p := color.New()
	typeCount := 0

	for idx, v := range n.Acquires {
		if v == 0 {
			continue
		}
		typeCount++
		p.Add(ColorConfig[idx])
		if idx == TokenIdxOnyx {
			require += p.Sprintf("%s", color.WhiteString("%d", v))
		} else {
			require += p.Sprintf("%d", v)
		}
	}

	// 前面补充空格，保持一致
	for i := 0; i < (4 - typeCount); i++ {
		require = " " + require
	}

	return fmt.Sprintf("[%d  %-4s]", n.Prestige, require)
}

func (s NobleTilesStack) Visual() string {
	info := ""
	for _, n := range s {
		info += n.Visual()
	}
	return info
}
