package service

import (
	"errors"
	"fmt"
	"splendor/model"
)

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

// PurchaseDevelopmentCard 玩家购买场上的发展卡
func PurchaseDevelopmentCard(p *model.Player, t *model.Table, cardIdx int) error {

	// 先判断场上是否有这个牌
	if !t.IsExistRevealedDevelopmentCard(cardIdx) {
		return errors.New(fmt.Sprintf("场上没这个牌，目标 %d，存在的牌 %+v", cardIdx, t.RevealedDevelopmentCards.ShowIdxInfo()))
	}

	// 移除场上的牌
	card, cardLevel, ok := t.RevealedDevelopmentCards.TakeCard(10001)
	if !ok {
		return errors.New(fmt.Sprintf("没有拿到卡牌啊，咋回事"))
	}

	// 补充场上的牌
	err := t.ReplaceRevealedDevelopmentCard(cardLevel)
	if err != nil {
		return err
	}

	// 给玩家添加卡牌
	p.AddDevelopmentCard(card)

	return nil
}
