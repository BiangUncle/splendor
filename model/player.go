package model

import (
	"errors"
	"fmt"
)

const HandCardUpperBound = 3

type Player struct {
	Name             string               // 玩家名字
	Tokens           TokenStack           // 宝石列表
	Bonuses          TokenStack           // 奖励列表
	DevelopmentCards DevelopmentCardStack // 发展卡列表
	HandCards        DevelopmentCardStack // 手中的发展卡
	NobleTitles      NobleTilesStack      // 贵族
	Prestige         int                  // 声望
}

// CreateANewPlayer 创建一个玩家
func CreateANewPlayer() *Player {
	return &Player{
		Tokens:           CreateEmptyTokenStack(),
		Bonuses:          CreateEmptyTokenStack(),
		DevelopmentCards: make(DevelopmentCardStack, 0),
		HandCards:        make(DevelopmentCardStack, 0),
		NobleTitles:      make(NobleTilesStack, 0),
		Prestige:         0,
	}
}

// AddTokens 玩家获取宝石
func (p *Player) AddTokens(tokens TokenStack) {
	p.Tokens.Add(tokens)
	return
}

// AddDevelopmentCard 玩家获取发展卡
func (p *Player) AddDevelopmentCard(card *DevelopmentCard) {
	// 发展卡增加这个
	p.DevelopmentCards = append(p.DevelopmentCards, card)
	p.Prestige += card.Prestige
	p.Bonuses[card.BonusType]++

	return
}

// AddNobleTile 玩家招待贵族
func (p *Player) AddNobleTile(noble *NobleTile) {
	// 发展卡增加这个
	p.NobleTitles = append(p.NobleTitles, noble)
	p.Prestige += noble.Prestige

	return
}

// HasEnoughToken 判断是否足够宝石
func (p *Player) HasEnoughToken(tokens TokenStack) bool {
	// 使用的黄金数量
	remainGoldJoker := p.Tokens[TokenIdxGoldJoker]

	for idx, tokensNum := range p.Tokens {
		// 不判断黄金
		if idx == TokenIdxGoldJoker {
			continue
		}
		// 判断现金加奖励够不够购买
		if tokensNum+p.Bonuses[idx] >= tokens[idx] {
			continue
		}
		// 判断加上黄金够不够，足够需要扣除黄金
		if tokensNum+p.Bonuses[idx]+remainGoldJoker >= tokens[idx] {
			remainGoldJoker -= tokens[idx] - (tokensNum + p.Bonuses[idx])
			continue
		}
		return false
	}

	return true
}

// PayToken 支付宝石
func (p *Player) PayToken(tokens TokenStack) (TokenStack, error) {
	// 使用的黄金数量
	remainGoldJoker := p.Tokens[TokenIdxGoldJoker]

	returnToken := CreateEmptyTokenStack()

	for idx, tokensNum := range p.Tokens {
		// 不判断黄金
		if idx == TokenIdxGoldJoker {
			continue
		}
		// 判断现金加奖励够不够购买
		if tokensNum+p.Bonuses[idx] >= tokens[idx] {
			needPay := tokens[idx] - p.Bonuses[idx] // 计算需要买多少
			returnToken[idx] += needPay
			continue
		}
		// 判断加上黄金够不够，足够需要扣除黄金
		if tokensNum+p.Bonuses[idx]+remainGoldJoker >= tokens[idx] {
			remainGoldJoker -= tokens[idx] - (tokensNum + p.Bonuses[idx])
			returnToken[idx] += tokensNum
			continue
		}
		return nil, errors.New(fmt.Sprintf("支付错误，根本不够啊，需要 %d, 只有 %d, 剩余黄金 %d。", tokens[idx], tokensNum+p.Bonuses[idx], remainGoldJoker))
	}

	// 计算需要返还的黄金
	returnToken[TokenIdxGoldJoker] = p.Tokens[TokenIdxGoldJoker] - remainGoldJoker

	// 角色扣除宝石
	err := p.Tokens.Minus(returnToken)
	if err != nil {
		return nil, err
	}

	return returnToken, nil
}

// AddHandCard 玩家获取手牌
func (p *Player) AddHandCard(card *DevelopmentCard) error {
	if len(p.HandCards) >= HandCardUpperBound {
		return errors.New(fmt.Sprintf("不能再增加手牌了，目前已经有 %d 张。", len(p.HandCards)))
	}
	// 发展卡增加这个
	p.HandCards = append(p.HandCards, card)
	return nil
}

// RemoveHandCard 移除手牌
func (p *Player) RemoveHandCard(cardIdx int) (*DevelopmentCard, error) {

	selectedIdx := -1
	var newHandCards []*DevelopmentCard

	for idx, card := range p.HandCards {
		if card.Idx == cardIdx {
			selectedIdx = idx
		} else {
			newHandCards = append(newHandCards, card)
		}
	}

	if selectedIdx == -1 {
		return nil, errors.New(fmt.Sprintf("移除手牌失败，cardIdx = %d, 现有手牌 %+v", cardIdx, p.HandCards))
	}

	ret := p.HandCards[selectedIdx]
	p.HandCards = newHandCards

	return ret, nil
}

// ReceiveNoble 招待贵族，如果不可以招待，返回 false
func (p *Player) ReceiveNoble(noble *NobleTile) (bool, error) {

	// 计算当前发展卡是否足够招待
	//existTokens := p.DevelopmentCards.ToTokenStack()
	existTokens := p.Bonuses // bonuses 和发展卡一个数量
	more := existTokens.MoreThan(noble.Acquires)
	if !more {
		return false, nil
	}

	// 将当前贵族加入自己名下
	p.AddNobleTile(noble)

	return true, nil
}

// ShowPlayerInfo 展示信息
func (p *Player) ShowPlayerInfo() {
	fmt.Printf("|==========Player==========\n")
	fmt.Printf("| Token:     %+v\n", p.Tokens)
	fmt.Printf("| Bonuses:   %+v\n", p.Bonuses)
	fmt.Printf("| Cards:     %+v\n", p.DevelopmentCards.ShowIdxInfo())
	fmt.Printf("| HandCards: %+v\n", p.HandCards.ShowIdxInfo())
	fmt.Printf("| Noble:     %+v\n", p.NobleTitles.ShowIdxInfo())
	fmt.Printf("| Prestige:  %+v\n", p.Prestige)
	fmt.Printf("|==========================\n")
}

// ShowPlayerInfoV2 展示信息
func (p *Player) ShowPlayerInfoV2() {

	infos := p.PlayerInfoString()

	fmt.Printf("|%s%-10s%s\n", "==========", " Player", "==========")
	for _, info := range infos {
		fmt.Printf("| %-30s\n", info)
	}
	fmt.Printf("|==============================\n")
}

// PlayerInfoString 玩家的信息
func (p *Player) PlayerInfoString() []string {
	return []string{
		fmt.Sprintf("%-10s %+v", "Token:", p.Tokens),
		fmt.Sprintf("%-10s %+v", "Bonuses:", p.Bonuses),
		fmt.Sprintf("%-10s %+v", "Cards:", p.DevelopmentCards.ShowIdxInfo()),
		fmt.Sprintf("%-10s %+v", "HandCards:", p.HandCards.ShowIdxInfo()),
		fmt.Sprintf("%-10s %+v", "Noble:", p.NobleTitles.ShowIdxInfo()),
		fmt.Sprintf("%-10s %+v", "Prestige:", p.Prestige),
	}
}
