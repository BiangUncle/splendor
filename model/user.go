package model

import "fmt"

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

func (p *Player) ShowPlayerInfo() {
	fmt.Printf("|=======================\n")
	fmt.Printf("| %+v\n", p.Tokens)
	fmt.Printf("| %+v\n", p.Bonuses)
	fmt.Printf("| %+v\n", p.DevelopmentCards.ShowIdxInfo())
	fmt.Printf("| %+v\n", p.NobleTitles)
	fmt.Printf("| %+v\n", p.Prestige)
	fmt.Printf("|=======================\n")
}
