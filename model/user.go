package model

import (
	"errors"
	"fmt"
)

type Player struct {
	Tokens           TokenStack           // 宝石列表
	Bonuses          TokenStack           // 奖励列表
	DevelopmentCards DevelopmentCardStack // 发展卡列表
	NobleTitles      NobleTilesStack      // 贵族
	Prestige         int                  // 声望
}

// CreateANewPlayer 创建一个玩家
func CreateANewPlayer() *Player {
	return &Player{
		Tokens:           CreateEmptyTokenStack(),
		Bonuses:          CreateEmptyTokenStack(),
		DevelopmentCards: make(DevelopmentCardStack, 0),
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

func (p *Player) ShowPlayerInfo() {
	fmt.Printf("|========Player=========\n")
	fmt.Printf("| Token: %+v\n", p.Tokens)
	fmt.Printf("| Bonuses: %+v\n", p.Bonuses)
	fmt.Printf("| Cards: %+v\n", p.DevelopmentCards.ShowIdxInfo())
	fmt.Printf("| Noble: %+v\n", p.NobleTitles)
	fmt.Printf("| Prestige: %+v\n", p.Prestige)
	fmt.Printf("|=======================\n")
}
