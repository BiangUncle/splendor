package model

import (
	"errors"
	"fmt"
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
