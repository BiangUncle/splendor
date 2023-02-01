package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"splendor/utils"
)

func Start() {

	InitDefaultAction()

	c := ConstructClient()
	g := ConstructGameStatus(c)

	for {
		fmt.Println(ShowOptionsInfos())
		action, err := utils.InputInt("")
		if err != nil {
			fmt.Println(err)
			continue
		}
		LoopAction(action)(g)
		if action == 0 {
			fmt.Println("exit")
			break
		}
	}

	return

	c = ConstructClient()
	g = ConstructGameStatus(c)

	for {
		g.ShowPlayerInfo()
		fmt.Print("请选择需要的操作 1.[加入] 2.[离开] 3.[探测] 4.[房间信息] 5.[下一位] 6.[拿三个宝石] 7.[拿两个宝石] 请选择:  ")

		action, err := utils.InputInt("")
		if err != nil {
			fmt.Println(err)
			break
		}

		switch action {
		case 3:
			content, err := g.Alive()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(content)
		case 5:
			if !g.IsOurTurn() {
				break
			}
			content, err := g.NextTurn()
			if err != nil {
				fmt.Println(err)
				break
			}
			nextPlayerName := gjson.Get(content, "current_player_name").String()
			fmt.Println(fmt.Sprintf("下一个是 %+v 操作", nextPlayerName))
		}
	}
}

func main() {
	Start()
}
