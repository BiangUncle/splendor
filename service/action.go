package service

import (
	"errors"
	"fmt"
	"splendor/model"
)

// TurnRound 玩家轮训
func TurnRound(table *model.Table) error {
	players := table.Players

	round := 0
	gameOver := false

	for !gameOver {
		for _, player := range players {
			// 执行动作
			err := TurnAction(player, table, 0)
			if err != nil {
				return err
			}
			// 招待贵族
			err = ReceiveNoble(player, table)
			if err != nil {
				return err
			}
			// 判断分数
			if player.Prestige >= 15 {
				gameOver = true
			}
		}
		round++
	}

	return nil
}

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

	// 判断是否有足够的宝石
	if !p.HasEnoughToken(model.DevelopmentCardMap[cardIdx].Acquires) {
		return errors.New(fmt.Sprintf("不够钱买啊，你需要 %+v，你只有 %+v%+v。", model.DevelopmentCardMap[cardIdx].Acquires, p.Tokens, p.Bonuses))
	}

	// 支付宝石
	returnToken, err := p.PayToken(model.DevelopmentCardMap[cardIdx].Acquires)
	if err != nil {
		return err
	}

	// 把返还的宝石返回到场上
	t.TokenStack.Add(returnToken)

	// 移除场上的牌
	card, cardLevel, ok := t.RevealedDevelopmentCards.TakeCard(10001) // todo: 拿走一张牌
	if !ok {
		return errors.New(fmt.Sprintf("没有拿到卡牌啊，咋回事"))
	}

	// 补充场上的牌
	err = t.ReplaceRevealedDevelopmentCard(cardLevel)
	if err != nil {
		return err
	}

	// 给玩家添加卡牌
	p.AddDevelopmentCard(card)

	return nil
}

// ReserveDevelopmentCard 保留一张发展卡
func ReserveDevelopmentCard(p *model.Player, t *model.Table, cardIdx int) error {

	// 判断是否手牌上限
	if len(p.HandCards) >= model.HandCardUpperBound {
		return errors.New(fmt.Sprintf("达到手牌上限了，加不了手牌了。"))
	}

	// 先判断场上是否有这个牌
	if !t.IsExistRevealedDevelopmentCard(cardIdx) {
		return errors.New(fmt.Sprintf("场上没这个牌，目标 %d，存在的牌 %+v", cardIdx, t.RevealedDevelopmentCards.ShowIdxInfo()))
	}

	// 拿走想要的一张牌
	card, cardLevel, ok := t.RevealedDevelopmentCards.TakeCard(10001) // todo: 拿走一张牌
	if !ok {
		return errors.New(fmt.Sprintf("没有拿到卡牌啊，咋回事"))
	}

	// 补充场上的牌
	err := t.ReplaceRevealedDevelopmentCard(cardLevel)
	if err != nil {
		return err
	}

	// 判断是否有黄金拿
	ok = t.TokenStack.TakeAGoldJoker()
	if ok {
		// 如果有黄金，就加入黄金
		p.Tokens.Add([]int{0, 0, 0, 0, 0, 1})
	}

	// 增加手牌
	err = p.AddHandCard(card)
	if err != nil {
		return err
	}

	return nil
}

// ReserveStackCard 从牌堆里面拿一张手牌
func ReserveStackCard(p *model.Player, t *model.Table) error {

	// 判断是否手牌上限
	if len(p.HandCards) >= model.HandCardUpperBound {
		return errors.New(fmt.Sprintf("达到手牌上限了，加不了手牌了。"))
	}

	// 从牌堆里面拿牌
	card, err := t.DevelopmentCardStacks.BottomStack.TakeTopCard()
	if err != nil {
		return err
	}

	// 判断是否有黄金拿
	ok := t.TokenStack.TakeAGoldJoker()
	if ok {
		// 如果有黄金，就加入黄金
		p.Tokens.Add([]int{0, 0, 0, 0, 0, 1})
	}

	// 增加手牌
	err = p.AddHandCard(card)
	if err != nil {
		return err
	}

	return nil
}

// PurchaseHandCard 购买手牌
func PurchaseHandCard(p *model.Player, t *model.Table, cardIdx int) error {
	// 先判断手里有没有牌
	if !p.HandCards.IsExistCard(cardIdx) {
		return errors.New(fmt.Sprintf("手牌没有这个呀，目标 %d，现有 %+v", cardIdx, p.HandCards.ShowIdxInfo()))
	}

	// 判断是否有足够的宝石
	if !p.HasEnoughToken(model.DevelopmentCardMap[cardIdx].Acquires) {
		return errors.New(fmt.Sprintf("不够钱买啊，你需要 %+v，你只有 %+v%+v。", model.DevelopmentCardMap[cardIdx].Acquires, p.Tokens, p.Bonuses))
	}

	// 支付宝石
	returnToken, err := p.PayToken(model.DevelopmentCardMap[cardIdx].Acquires)
	if err != nil {
		return err
	}

	// 把返还的宝石返回到场上
	t.TokenStack.Add(returnToken)

	// 移除手牌
	card, err := p.RemoveHandCard(cardIdx)
	if err != nil {
		return err
	}

	// 转换为自己的发展卡
	p.AddDevelopmentCard(card)

	return nil
}

// ReceiveNoble 招待贵族
func ReceiveNoble(p *model.Player, t *model.Table) error {

	for idx, noble := range t.RevealedNobleTiles {
		// 没有贵族了
		if noble == nil {
			continue
		}

		// 判断是否可以招待贵族
		ok, err := p.ReceiveNoble(noble)
		if err != nil {
			return err
		}
		// 无法招待，继续判断
		if !ok {
			continue
		}

		// 移除贵族
		t.RevealedNobleTiles[idx] = nil
		break
	}

	return nil
}
