package model

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
)

const TokenTypeNumber = 6

// todo 后续可配置
var (
	EmeraldTokenNumber   = 7 // 翡翠的
	DiamondTokenNumber   = 7 // 金刚石
	SapphireTokenNumber  = 7 // 蓝宝石
	OnyxTokenNumber      = 7 // 玛瑙
	RubyTokenNumber      = 7 // 红宝石
	GoldJokerTokenNumber = 5 // 黄金
)

const (
	TokenIdxEmerald = iota
	TokenIdxDiamond
	TokenIdxSapphire
	TokenIdxOnyx
	TokenIdxRuby
	TokenIdxGoldJoker
)

var ColorConfig = map[int]color.Attribute{
	TokenIdxEmerald:   color.BgGreen,
	TokenIdxDiamond:   color.BgWhite,
	TokenIdxSapphire:  color.BgBlue,
	TokenIdxOnyx:      color.BgHiMagenta,
	TokenIdxRuby:      color.BgRed,
	TokenIdxGoldJoker: color.BgYellow,
}

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

// CreatANewTokenStack 创建一个新的桌面宝石堆
func CreatANewTokenStack() TokenStack {
	return defaultTokenStack.Copy()
}

// CreateEmptyTokenStack 创建一个空白的宝石堆数组
func CreateEmptyTokenStack() TokenStack {
	return make(TokenStack, TokenTypeNumber)
}

// IntList2TokenStack 将int列表转tokenStack
func IntList2TokenStack(i []int) (TokenStack, error) {
	selected := make(TokenStack, TokenTypeNumber)

	for _, typ := range i {
		if typ == -1 {
			continue
		}
		if typ >= TokenTypeNumber || typ < 0 {
			return nil, errors.New(fmt.Sprintf("invalid token index: %+v", typ))
		}
		selected[typ]++
	}
	return selected, nil
}

// Strings2TokenStack token字符串转对象
func Strings2TokenStack(tokensStrings string) (TokenStack, error) {
	var tokens []int
	for i := 0; i < len(tokensStrings); i++ {
		tokens = append(tokens, int(tokensStrings[i]-'0'))
	}
	return IntList2TokenStack(tokens)
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

// TakeThreeTokens 拿取三个不同宝石
func (s TokenStack) TakeThreeTokens(tokens []int) (TokenStack, error) {

	if len(tokens) > 3 {
		return nil, errors.New(fmt.Sprintf("目前不支持拿大于3数量的宝石，只能拿3，你拿了 %d", len(tokens)))
	}

	// 判断是否拿得到
	selected := make(TokenStack, TokenTypeNumber)

	for _, typ := range tokens {
		if typ == -1 {
			continue
		}
		selected[typ]++
	}

	return s.TakeTokens(selected)
}

// TakeDoubleTokens 拿取两个不同宝石
func (s TokenStack) TakeDoubleTokens(tokenIdx int) (TokenStack, error) {
	// 判断是否拿得到
	selected := make(TokenStack, TokenTypeNumber)

	selected[tokenIdx] = 2

	return s.TakeTokens(selected)
}

// TakeTokens 拿去一定数量的宝石，对于3/2个宝石的拿取函数的一个抽象
func (s TokenStack) TakeTokens(selected TokenStack) (TokenStack, error) {
	// 判断是否拿得到
	ret := make(TokenStack, TokenTypeNumber)
	var err error
	var rollback []func()

	defer func() {
		if err != nil {
			for _, rb := range rollback {
				rb()
			}
		}
	}()

	for idx, need := range selected {
		if need == 0 {
			continue
		}
		if s[idx] < need {
			err = errors.New(fmt.Sprintf("不够拿宝石，需要 %d 个，只有 %d 个。", need, s[idx]))
			return nil, err
		}
		ret[idx] += need
		// 执行操作
		s[idx] -= need
		// 回滚操作
		rollback = append(rollback, func() { s[idx] += need })
	}

	return ret, nil
}

// TakeAGoldJoker 拿走一个黄金
func (s TokenStack) TakeAGoldJoker() bool {
	if s[TokenIdxGoldJoker] < 1 {
		return false
	}
	s[TokenIdxGoldJoker]--
	return true
}

// Add 添加宝石
func (s TokenStack) Add(tokens TokenStack) {
	for idx, v := range tokens {
		s[idx] += v
	}
	return
}

// Minus 减少宝石
func (s TokenStack) Minus(tokens TokenStack) error {

	var err error
	var rollback []func()

	defer func() {
		if err != nil {
			for _, rb := range rollback {
				rb()
			}
		}
	}()

	for idx, v := range tokens {
		if s[idx] < v {
			return errors.New(fmt.Sprintf("无法扣除宝石，只有 %d 个，要扣 %d 个。", s[idx], v))
		}
		s[idx] -= v
		rollback = append(rollback, func() {
			s[idx] += v
		})
	}
	return nil
}

// MoreThan 判断是否所有宝石都多过
func (s TokenStack) MoreThan(tokens TokenStack) bool {
	for idx, v := range s {
		if v < tokens[idx] {
			return false
		}
	}
	return true
}

// Count 计算宝石数量
func (s TokenStack) Count() int {
	count := 0
	for _, n := range s {
		count += n
	}
	return count
}

// Copy 复制宝石列表
func (s TokenStack) Copy() TokenStack {
	cpy := make(TokenStack, len(s))
	copy(cpy, s)
	return cpy
}
