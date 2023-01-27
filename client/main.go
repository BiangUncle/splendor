package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"splendor/utils"
)

func main() {
	c := ConstructClient()

	g := &GameStatus{
		Client: c,
	}

	for {
		fmt.Println(g.Info())
		fmt.Print("请选择需要的操作 1. 加入 2. 离开 3. 探测 4. 房间信息: ")

		action, err := utils.InputInt()
		if err != nil {
			fmt.Println(err)
			break
		}
		switch action {
		case 1:
			content, err := c.JoinGame()
			if err != nil {
				fmt.Println(err)
				break
			}
			g.ConnectStatus = gjson.Get(content, "status").String()
			g.TableID = gjson.Get(content, "table_id").String()
			g.PlayerID = gjson.Get(content, "player_id").String()
			g.SessionID = gjson.Get(content, "session_id").String()
			g.UserName = gjson.Get(content, "username").String()

		case 2:
			content, err := c.Leave()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(content)
		case 3:
			content, err := c.Alive()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(content)
		case 4:
			content, err := c.TableInfo()
			if err != nil {
				fmt.Println(err)
				break
			}
			tableInfo := gjson.Get(content, "tableInfo").String()
			fmt.Println(tableInfo)
		case 0:
			content, err := c.Leave()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("退出游戏")
			fmt.Println(content)
			return
		}
	}

}
