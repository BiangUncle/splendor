package model

import (
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

type NobleTilesStack []NobleTile

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

func (s NobleTilesStack) ShowIdxInfo() {
	idxInfo := make([]int, len(s))
	for i, noble := range s {
		idxInfo[i] = noble.Idx
	}
	fmt.Printf("%+v\n", idxInfo)
}
