package service

import (
	"fmt"
	"splendor/model"
	"splendor/utils"
)

// CreateANewGame 创建一局新的游戏
func CreateANewGame(playerNum int) (*model.Table, error) {
	table := model.CreateTable()

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

// TurnRound 玩家轮训
func TurnRound(table *model.Table) error {
	players := table.Players

	round := 0
	gameOver := false

	for !gameOver {
		for _, player := range players {
			// 执行动作
			err := ActionCMD(player, table)
			if err != nil {
				fmt.Println(err)
				continue
			}
			// 招待贵族
			err = ReceiveNoble(player, table)
			if err != nil {
				fmt.Println(err)
				continue
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

// Action 玩家进行动作
func Action(p *model.Player, t *model.Table, action int) error {
	switch action {
	case 1: // 抽三个宝石
		fmt.Print("请抽取三个不同的宝石: ")
		tokenIdx, err := utils.InputList()
		if err != nil {
			return err
		}
		return ActionTakeThreeTokens(p, t, tokenIdx)
	case 2: // 抽两个一样的宝石
		fmt.Print("请抽取两个相同的宝石: ")
		tokenId, err := utils.InputInt()
		if err != nil {
			return err
		}
		return ActionTakeDoubleTokens(p, t, tokenId)
	case 3: // 购买一张发展牌
		cardIdx, err := utils.InputInt()
		if err != nil {
			return err
		}
		err = PurchaseDevelopmentCard(p, t, cardIdx)
		if err != nil {
			return err
		}
		break
	case 4: // 预购/摸取一张发展牌
		cardIdx, err := utils.InputInt()
		if err != nil {
			return err
		}
		err = ReserveDevelopmentCard(p, t, cardIdx)
		if err != nil {
			return err
		}
		break
	}
	return nil
}
