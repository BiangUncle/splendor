package service

import (
	"fmt"
	"splendor/model"
)

// CreateANewGame 创建一局新的游戏
func CreateANewGame(playerNum int) (*model.Table, error) {
	table := model.CreateANewTable()

	for i := 0; i < playerNum; i++ {
		player := model.CreateANewPlayer()
		player.Name = fmt.Sprintf("玩家%d", i+1)
		table.AddPlayer(player)
	}

	err := table.Reveal()
	if err != nil {
		return nil, err
	}

	table.ShowTableInfo()
	return table, nil
}

// TurnRoundCMD 玩家轮训
func TurnRoundCMD(table *model.Table) error {
	players := table.Players

	round := 0
	gameOver := false

	for !gameOver {
		for _, player := range players {
			// 执行动作
			err := ActionCMD(player, table)
			if err != nil {
				fmt.Println(err)
				return err
			}
			// 招待贵族
			err = ReceiveNoble(player, table)
			if err != nil {
				fmt.Println(err)
				return err
			}
			// 判断分数
			if player.Prestige >= 15 {
				gameOver = true
			}
		}
		round++
		table.ShowTableInfo()
	}

	return nil
}

// ActionCMD 命令行判断执行逻辑
func ActionCMD(p *model.Player, t *model.Table) (err error) {
	var action int
	fmt.Printf("[%s]需要什么操作: ", p.Name)
	_, err = fmt.Scanf("%d", &action)
	if err != nil {
		return
	}
	err = Action(p, t, action)
	return
}
