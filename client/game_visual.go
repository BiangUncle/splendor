package main

import (
	"fmt"
	"github.com/fatih/color"
	"splendor/utils"
)

func (g *GameStatus) Info() []string {

	c := color.New(color.FgCyan)

	return []string{
		fmt.Sprintf("玩家[%6s] ", c.Sprint(g.UserName)),
		fmt.Sprintf("状态[%6s] ", c.Sprint(g.ConnectStatus)),
		fmt.Sprintf("会话[%6s] ", c.Sprint(utils.CompressUuid(g.SessionID))),
		fmt.Sprintf("房间[%6s] ", c.Sprint(utils.CompressUuid(g.TableID))),
		fmt.Sprintf("玩家[%6s] ", c.Sprint(utils.CompressUuid(g.PlayerID))),
	}
}

func (g *GameStatus) ShowPlayerInfo() string {
	ret := ""
	ret += "\033[40m                                                                                   \033[0m\n"
	for _, info := range g.Info() {
		ret += info
	}
	ret += "\n"
	ret += "\033[40m                                                                                   \033[0m"
	return ret
}

func (g *GameStatus) ReturnContent() string {
	ret := ""
	ret += g.RetContent + "\n"
	ret += "\033[40m                                                                                   \033[0m\n"
	return ret
}
