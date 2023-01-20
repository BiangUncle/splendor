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

func (p *Player) AddTokens(tokens TokenStack) {
	p.Tokens.Add(tokens)
	return
}

func (p *Player) ShowPlayerInfo() {
	fmt.Printf("|=======================\n")
	fmt.Printf("| %+v\n", p.Tokens)
	fmt.Printf("| %+v\n", p.Bonuses)
	fmt.Printf("| %+v\n", p.DevelopmentCards)
	fmt.Printf("| %+v\n", p.NobleTitles)
	fmt.Printf("| %+v\n", p.Prestige)
	fmt.Printf("|=======================\n")

}
