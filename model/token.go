package model

import (
	"errors"
	"fmt"
	"math/rand"
)

const TokenTypeNumber = 6

// todo 后续可配置
var (
	EmeraldTokenNumber   = 7
	DiamondTokenNumber   = 7
	SapphireTokenNumber  = 7
	OnyxTokenNumber      = 7
	RubyTokenNumber      = 7
	GoldJokerTokenNumber = 5
)

const (
	TokenIdxEmerald = iota
	TokenIdxDiamond
	TokenIdxSapphire
	TokenIdxOnyx
	TokenIdxRuby
	TokenIdxGoldJoker
)

type Token struct {
}

var defaultTokenStack TokenStack

type TokenStack []int

// InitTokenStack 初始化默认宝石堆
func InitTokenStack() {
	defaultTokenStack = make(TokenStack, 6)

	defaultTokenStack[TokenIdxEmerald] = EmeraldTokenNumber
	defaultTokenStack[TokenIdxDiamond] = DiamondTokenNumber
	defaultTokenStack[TokenIdxSapphire] = SapphireTokenNumber
	defaultTokenStack[TokenIdxOnyx] = OnyxTokenNumber
	defaultTokenStack[TokenIdxRuby] = RubyTokenNumber
	defaultTokenStack[TokenIdxGoldJoker] = GoldJokerTokenNumber
}

// takeToken 拿走一定数量的宝石
func (s TokenStack) takeToken(tokenId int, num int) (TokenStack, error) {
	// 判断是否拿得到
	if s[tokenId] < num {
		return TokenStack{}, errors.New(fmt.Sprintf("不够拿宝石，需要 %d 个，只有 %d 个。", num, s[tokenId]))
	}

	ret := TokenStack{}
	s[tokenId] -= num
	ret[tokenId] += num

	return ret, nil
}

func (s TokenStack) TakeThreeTokens(tokens TokenStack) (TokenStack, error) {
	// 判断是否拿得到
	ret := TokenStack{}

	for idx, v := range tokens {
		if s[idx] < v {
			return nil, errors.New(fmt.Sprintf("不够拿宝石，需要 %d 个，只有 %d 个。", v, s[idx]))
		}
		ret[idx]++
		s[idx]--
	}

	return ret, nil
}

// Add 添加宝石
func (s TokenStack) Add(tokens TokenStack) {
	for idx, v := range tokens {
		s[idx] += v
	}
	return
}

func CreatANewTokenStack() TokenStack {
	return defaultTokenStack.Copy()
}

func (s TokenStack) Copy() TokenStack {
	cpy := make(TokenStack, len(s))
	copy(cpy, s)
	return cpy
}

func ShuffleToken(tokens []Token) []Token {
	rand.Shuffle(len(tokens), func(i, j int) {
		tokens[i], tokens[j] = tokens[j], tokens[i]
	})
	return tokens
}
