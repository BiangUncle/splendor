package service

import "splendor/model"

func TurnAction(u *model.User, t *model.Table, action int) error {
	switch action {
	case 1: // 抽三个宝石
		tokenStack, err := t.TokenStack.TakeThreeTokens([]int{1, 1, 1, 0, 0, 0})
		if err != nil {
			return err
		}
		u.AddTokens(tokenStack)
	case 2:
	// 抽两个一样的宝石
	case 3:
	// 购买一张发展牌
	case 4: // 预购/摸取一张发展牌
	}
	return nil
}
