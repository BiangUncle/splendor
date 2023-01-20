package model

import "math/rand"

const NobleTitleNumber = 10

type NobleTile struct {
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
func (n NobleTilesStack) Copy() NobleTilesStack {
	cpy := make(NobleTilesStack, len(n))
	copy(cpy, n)
	return cpy
}

func ShuffleNobleTitle(nobleTitles []NobleTile) []NobleTile {
	rand.Shuffle(len(nobleTitles), func(i, j int) {
		nobleTitles[i], nobleTitles[j] = nobleTitles[j], nobleTitles[i]
	})
	return nobleTitles
}
