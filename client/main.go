package main

import (
	"fmt"
	"splendor/utils"
)

func main() {
	c := ConstructClient()

	g := &GameStatus{
		Client: c,
	}

	for {
		fmt.Println(fmt.Sprintf("[GAME]状态: %+v", g.ConnectStatus))
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
			fmt.Println(content)
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
			fmt.Println(content)
		case 0:
			fmt.Println("退出游戏")
			return
		}
	}

}
