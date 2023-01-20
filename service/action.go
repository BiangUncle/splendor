package service

import "splendor/model"

// TurnAction 玩家进行动作
func TurnAction(p *model.Player, t *model.Table, action int) error {
	switch action {
	case 1: // 抽三个宝石
		tokenStack, err := t.TokenStack.TakeThreeTokens([]int{1, 1, 1, 0, 0, 0}) // todo: 用户选择
		if err != nil {
			return err
		}
		p.AddTokens(tokenStack)
	case 2: // 抽两个一样的宝石
		tokenStack, err := t.TokenStack.TakeDoubleTokens(model.TokenIdxEmerald) // todo: 用户选择
		if err != nil {
			return err
		}
		p.AddTokens(tokenStack)
	case 3:
	// 购买一张发展牌
	case 4: // 预购/摸取一张发展牌
	}
	return nil
}
