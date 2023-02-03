package model

import (
	"errors"
	"fmt"
	"splendor/utils"
)

/*
这个文件主要负责模组之间的调用关系
*/

// PurchaseDevelopmentCard 买卡操作
func PurchaseDevelopmentCard(p *Player, t TokenStack, table *Table, cardIdx int) error {

	// 先把钱拿出来
	tokens, err := p.TakeOutTokens(t)
	if err != nil {
		return errors.New(fmt.Sprintf("你拿不出钱啊，你需要 %+v，你只有 %+v。", tokens, p.Tokens))
	}
	returnToken := tokens.Copy()

	// token加上bonus
	tokens.Add(p.Bonuses)

	if !tokens.MoreThan(DevelopmentCardMap[cardIdx].Acquires) {
		return errors.New(fmt.Sprintf("你不够钱啊，你需要 %+v，加上奖金你也只有 %+v。", DevelopmentCardMap[cardIdx].Acquires, p.Tokens))
	}

	// 把返还的宝石返回到场上
	table.TokenStack.Add(returnToken)

	// 移除场上的牌
	card, cardLevel, ok := table.RevealedDevelopmentCards.TakeCard(cardIdx)
	if !ok {
		return errors.New(fmt.Sprintf("没有拿到卡牌啊，咋回事"))
	}

	// 补充场上的牌
	err = table.ReplaceRevealedDevelopmentCard(cardLevel)
	if err != nil {
		return err
	}

	// 给玩家添加卡牌
	p.AddDevelopmentCard(card)

	return nil
}

// PurchaseHandCard 买卡操作
func PurchaseHandCard(p *Player, t TokenStack, table *Table, cardIdx int) error {

	// 先把钱拿出来
	tokens, err := p.TakeOutTokens(t)
	if err != nil {
		return errors.New(fmt.Sprintf("你拿不出钱啊，你需要 %+v，你只有 %+v。", tokens, p.Tokens))
	}
	returnToken := tokens.Copy()

	// token加上bonus
	tokens.Add(p.Bonuses)

	if !tokens.MoreThan(DevelopmentCardMap[cardIdx].Acquires) {
		return errors.New(fmt.Sprintf("你不够钱啊，你需要 %+v，加上奖金你也只有 %+v。", DevelopmentCardMap[cardIdx].Acquires, p.Tokens))
	}

	// 把返还的宝石返回到场上
	table.TokenStack.Add(returnToken)

	// 移除手牌
	card, err := p.RemoveHandCard(cardIdx)
	if err != nil {
		return err
	}

	// 给玩家添加卡牌
	p.AddDevelopmentCard(card)

	return nil
}

// ReceiveNoble 招待贵族
func ReceiveNoble(p *Player, t *Table) error {

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
		err = t.RemoveRevealedNoble(idx)
		if err != nil {
			return err
		}
		break
	}

	return nil
}

// ActionTakeThreeTokens 执行拿3个不同宝石
func ActionTakeThreeTokens(p *Player, t *Table, tokenIdx []int) (int, error) {
	tokenStack, err := t.TokenStack.TakeThreeTokens(tokenIdx)
	if err != nil {
		return 0, err
	}
	p.AddTokens(tokenStack)
	ret := utils.Max(0, p.Tokens.Count()-TokensNumberUpperLimit) // 需要返还多少
	return ret, nil
}

// ActionTakeDoubleTokens 执行拿2个相同宝石
func ActionTakeDoubleTokens(p *Player, t *Table, tokenIdx int) (int, error) {
	tokenStack, err := t.TokenStack.TakeDoubleTokens(tokenIdx)
	if err != nil {
		return 0, err
	}
	p.AddTokens(tokenStack)

	ret := utils.Max(0, p.Tokens.Count()-TokensNumberUpperLimit) // 需要返还多少
	return ret, nil
}

// ReserveDevelopmentCard 保留一张发展卡
func ReserveDevelopmentCard(p *Player, t *Table, cardIdx int) error {

	// 判断是否手牌上限
	if len(p.HandCards) >= HandCardUpperBound {
		return errors.New(fmt.Sprintf("达到手牌上限了，加不了手牌了。"))
	}

	// 先判断场上是否有这个牌
	if !t.IsExistRevealedDevelopmentCard(cardIdx) {
		return errors.New(fmt.Sprintf("场上没这个牌，目标 %d，存在的牌 %+v", cardIdx, t.RevealedDevelopmentCards.ShowIdxInfo()))
	}

	// 拿走想要的一张牌
	card, cardLevel, ok := t.RevealedDevelopmentCards.TakeCard(cardIdx)
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

// ActionReturnTokens 角色返还多余的宝石
func ActionReturnTokens(p *Player, t *Table, tokens TokenStack) error {

	// 扣除角色身上的宝石
	err := p.Tokens.Minus(tokens)
	if err != nil {
		return err
	}
	// 将宝石返还给桌台
	t.TokenStack.Add(tokens)
	return nil
}
