package model

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
)

func (t *Table) ShowVisualInfo() {
	fmt.Println("\033[40m                                                                \033[0m")
	fmt.Println(t.TokenStack.Visual())
	fmt.Printf("[%d] %s\n", len(t.DevelopmentCardStacks.TopStack), t.RevealedDevelopmentCards.TopStack.Visual())
	fmt.Printf("[%d] %s\n", len(t.DevelopmentCardStacks.MiddleStack), t.RevealedDevelopmentCards.MiddleStack.Visual())
	fmt.Printf("[%d] %s\n", len(t.DevelopmentCardStacks.BottomStack), t.RevealedDevelopmentCards.BottomStack.Visual())
	fmt.Println(t.RevealedNobleTiles.Visual())
	fmt.Println("\033[40m                                                                \033[0m")
}

func (t *Table) ALLPlayerTabInfo() []string {
	var players []string
	for _, p := range t.Players {
		players = append(players, fmt.Sprintf("|%s|", p.Name))
	}
	return players
}

func (t *Table) PlayerTabInfo() string {
	players := t.ALLPlayerTabInfo()
	if len(players) == 0 {
		return ""
	}
	c := color.New(color.BgMagenta)
	ret := ""
	for idx, player := range players {
		if idx == t.CurrentPlayerIdx {
			ret += c.Sprint(player)
		} else {
			ret += player
		}
	}
	return ret
}

// TableInfoString 玩家的信息
func (t *Table) TableInfoString() []string {
	return []string{
		fmt.Sprintf("%-15s %+v", "Name:", t.Name),
		fmt.Sprintf("%-15s %+v", "CurPlayer:", t.PlayerTabInfo()),
		fmt.Sprintf("%-15s %+v", "Token:", t.TokenStack),
		fmt.Sprintf("%-15s %+v", "DevCard:", t.DevelopmentCardStacks.ShowIdxInfo()),
		fmt.Sprintf("%-15s %+v", "RevealCards:", t.RevealedDevelopmentCards.ShowIdxInfo()),
		fmt.Sprintf("%-15s %+v", "Noble:", t.NobleTilesStack.ShowIdxInfo()),
		fmt.Sprintf("%-15s %+v", "RevealNoble:", t.RevealedNobleTiles.ShowIdxInfo()),
	}
}

// ShowInfo 展示信息
func (t *Table) ShowInfo() {
	fmt.Printf("|=========Table=========\n")
	fmt.Printf("| Token: %+v\n", t.TokenStack)
	fmt.Printf("| DevCard: %+v\n", t.DevelopmentCardStacks.ShowIdxInfo())
	fmt.Printf("| RevealCards: %+v\n", t.RevealedDevelopmentCards.ShowIdxInfo())
	fmt.Printf("| Noble: %+v\n", t.NobleTilesStack.ShowIdxInfo())
	fmt.Printf("| RevealNoble: %+v\n", t.RevealedNobleTiles.ShowIdxInfo())
	fmt.Printf("|=======================\n")
}

// ShowTableInfo 展示整场游戏信息
func (t *Table) ShowTableInfo() string {

	infos := make([][]string, 0)

	ret := ""

	for _, player := range t.Players {
		infos = append(infos, player.PlayerInfoString())
	}

	line := ""
	for j := 0; j < len(infos); j++ {
		line = line + fmt.Sprintf("%s", "\u001B[40m                                \u001B[0m")
	}

	left := "\u001B[40m \u001B[0m"

	ret = ret + fmt.Sprintf("%s\n", line)

	if len(infos) > 0 {
		for i := 0; i < len(infos[0]); i++ {
			infoRow := ""
			for j := 0; j < len(infos); j++ {
				infoRow = infoRow + fmt.Sprintf("%s %-30s", left, infos[j][i])
			}
			infoRow = infoRow + "\n"
			ret = ret + infoRow
		}
	}

	ret = ret + fmt.Sprintf("%s\n", line)

	for _, info := range t.TableInfoString() {
		ret = ret + fmt.Sprintf("%s %s\n", left, info)
	}

	ret = ret + fmt.Sprintf("%s\n", line)

	return ret
}

func (t *Table) Marshal() (string, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (t *Table) Unmarshal(s string) error {
	err := json.Unmarshal([]byte(s), t)
	if err != nil {
		return err
	}
	return nil
}
