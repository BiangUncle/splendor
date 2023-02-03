package model

import "fmt"

// ShowPlayerInfo 展示信息
func (p *Player) ShowPlayerInfo() {
	fmt.Printf("|==========Player==========\n")
	fmt.Printf("| Name:      %+v\n", p.Name)
	fmt.Printf("| Token:     %+v\n", p.Tokens)
	fmt.Printf("| Bonuses:   %+v\n", p.Bonuses)
	fmt.Printf("| Cards:     %+v\n", p.DevelopmentCards.ShowIdxInfo())
	fmt.Printf("| HandCards: %+v\n", p.HandCards.ShowIdxInfo())
	fmt.Printf("| Noble:     %+v\n", p.NobleTitles.ShowIdxInfo())
	fmt.Printf("| Prestige:  %+v\n", p.Prestige)
	fmt.Printf("|==========================\n")
}

// PlayerInfoString 玩家的信息
func (p *Player) PlayerInfoString() []string {
	return []string{
		fmt.Sprintf("%-10s %+v", "Name:", p.Name),
		fmt.Sprintf("%-10s %+v", "Token:", p.Tokens),
		fmt.Sprintf("%-10s %+v", "Bonuses:", p.Bonuses),
		fmt.Sprintf("%-10s %+v", "Cards:", p.DevelopmentCards.ShowIdxInfo()),
		fmt.Sprintf("%-10s %+v", "HandCards:", p.HandCards.ShowIdxInfo()),
		fmt.Sprintf("%-10s %+v", "Noble:", p.NobleTitles.ShowIdxInfo()),
		fmt.Sprintf("%-10s %+v", "Prestige:", p.Prestige),
	}
}

//func (p *Player) WholeVisual() []string {
//	 var ret []string
//
//}