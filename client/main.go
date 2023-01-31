package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"splendor/utils"
)

func Start() {
	c := ConstructClient()

	g := ConstructGameStatus(c)
	g.wg.Add(1)

	for {
		g.ShowPlayerInfo()
		fmt.Print("请选择需要的操作 1.[加入] 2.[离开] 3.[探测] 4.[房间信息] 5.[下一位] 6.[拿三个宝石] 7.[拿两个宝石] 请选择:  ")

		action, err := utils.InputInt()
		if err != nil {
			fmt.Println(err)
			break
		}

		switch action {
		case 1:
			content, err := g.JoinGame()
			if err != nil {
				fmt.Println(err)
				break
			}
			g.ConnectStatus = gjson.Get(content, "status").String()
			g.TableID = gjson.Get(content, "table_id").String()
			g.PlayerID = gjson.Get(content, "player_id").String()
			g.SessionID = gjson.Get(content, "session_id").String()
			g.UserName = gjson.Get(content, "username").String()

			go g.RoutineKeepALive()
		case 2:
			content, err := g.Leave()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(content)
			g.Stop()
		case 3:
			content, err := g.Alive()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(content)
		case 4:
			content, err := g.TableInfo()
			if err != nil {
				fmt.Println(err)
				break
			}
			tableInfo := gjson.Get(content, "tableInfo").String()
			fmt.Println(tableInfo)
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
		case 6:
			if !g.IsOurTurn() {
				break
			}
			tokensString, err := utils.InputString()
			if err != nil {
				fmt.Println(err)
				break
			}
			msg, err := g.TakeThreeTokens(tokensString)
			if err != nil {
				fmt.Println(err)
				break
			}
			if msg == "ret" {
				_, err = g.ReturnTokens(inputString())
				if err != nil {
					fmt.Println(err)
					break
				}
			}
		case 7:
			if !g.IsOurTurn() {
				break
			}
			tokenId, err := utils.InputString()
			if err != nil {
				fmt.Println(err)
				break
			}
			msg, err := g.TakeDoubleTokens(utils.ToInt(tokenId))
			if err != nil {
				fmt.Println(err)
				break
			}
			if msg == "ret" {
				_, err = g.ReturnTokens(inputString())
				if err != nil {
					fmt.Println(err)
					break
				}
			}
		case 0:
			content, err := g.Leave()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("客户端: ", "退出游戏")
			fmt.Println("服务端: ", content)
			g.Stop()
			g.wg.Wait()
			return
		}
	}
}

func main() {
	Start()
}
