package model

import (
	"errors"
	"fmt"
	"time"
)

const RevealedDevelopmentCardNumPerLevel = 4

// Table 游戏桌面
type Table struct {
	Players  []*Player // 玩家
	GameTime time.Time // 游戏时间

	DevelopmentCardStacks    *DevelopmentCardStacks // 发展卡堆
	RevealedDevelopmentCards *DevelopmentCardStacks // 暴露的发展卡

	NobleTilesStack    NobleTilesStack // 贵族卡堆
	RevealedNobleTiles []*NobleTile    // 暴露的贵族卡

	TokenStack TokenStack // 宝石卡堆

	CurrentPlayer *Player // 当前角色
}

// CreateANewTable 开一桌
func CreateANewTable() *Table {
	table := &Table{
		Players:                  make([]*Player, 0),
		GameTime:                 time.Now(),
		DevelopmentCardStacks:    CreateANewDevelopmentCardStacks(),
		NobleTilesStack:          CreateANewNobleTilesStack(),
		TokenStack:               CreatANewTokenStack(),
		RevealedDevelopmentCards: CreateEmptyDevelopmentCardStacks(),
		RevealedNobleTiles:       make([]*NobleTile, 3),
	}

	return table
}

func (t *Table) AddPlayer(p *Player) {
	t.Players = append(t.Players, p)
}

func (t *Table) Shuffle() {
	t.DevelopmentCardStacks.Shuffle()
	t.NobleTilesStack.Shuffle()
}

// RevealDevelopmentCard 初始化发展卡
func (t *Table) RevealDevelopmentCard() (err error) {
	topCards, err := t.DevelopmentCardStacks.TopStack.TakeTopNCard(RevealedDevelopmentCardNumPerLevel)
	if err != nil {
		return
	}
	t.RevealedDevelopmentCards.TopStack = topCards

	middleCards, err := t.DevelopmentCardStacks.MiddleStack.TakeTopNCard(RevealedDevelopmentCardNumPerLevel)
	if err != nil {
		return
	}
	t.RevealedDevelopmentCards.MiddleStack = middleCards

	bottomCards, err := t.DevelopmentCardStacks.BottomStack.TakeTopNCard(RevealedDevelopmentCardNumPerLevel)
	if err != nil {
		return
	}
	t.RevealedDevelopmentCards.BottomStack = bottomCards

	return nil
}

// RevealNobleTiles 初始化贵族卡
func (t *Table) RevealNobleTiles() (err error) {

	userNum := len(t.Players)

	nobleTiles, err := t.NobleTilesStack.TakeTopNCard(userNum + 1)
	if err != nil {
		return
	}
	t.RevealedNobleTiles = nobleTiles

	return nil
}

// Reveal 展示发展卡和贵族卡
func (t *Table) Reveal() (err error) {
	err = t.RevealDevelopmentCard()
	if err != nil {
		return
	}
	err = t.RevealNobleTiles()
	if err != nil {
		return
	}
	return nil
}

// IsExistRevealedDevelopmentCard 判断这个卡是否在场上
func (t *Table) IsExistRevealedDevelopmentCard(idx int) bool {
	return t.RevealedDevelopmentCards.IsExistCard(idx)
}

// ReplaceRevealedDevelopmentCard 补充场上的发展卡
// todo: 可优化，有大量重复逻辑
func (t *Table) ReplaceRevealedDevelopmentCard(cardLevel int) error {

	switch cardLevel {
	case DevelopmentCardLevelTop:
		card, err := t.DevelopmentCardStacks.TopStack.TakeTopCard()
		// 如果没卡就不补充
		if err != nil && len(t.DevelopmentCardStacks.TopStack) == 0 {
			return nil
		}
		if err != nil {
			return err
		}
		err = t.RevealedDevelopmentCards.TopStack.PutNewCardToEmptySite(card)
		if err != nil {
			return err
		}
		return nil
	case DevelopmentCardLevelMiddle:
		card, err := t.DevelopmentCardStacks.MiddleStack.TakeTopCard()
		// 如果没卡就不补充
		if err != nil && len(t.DevelopmentCardStacks.MiddleStack) == 0 {
			return nil
		}
		if err != nil {
			return err
		}
		err = t.RevealedDevelopmentCards.MiddleStack.PutNewCardToEmptySite(card)
		if err != nil {
			return err
		}
		return nil
	case DevelopmentCardLevelBottom:
		card, err := t.DevelopmentCardStacks.BottomStack.TakeTopCard()
		// 如果没卡就不补充
		if err != nil && len(t.DevelopmentCardStacks.BottomStack) == 0 {
			return nil
		}
		if err != nil {
			return err
		}
		err = t.RevealedDevelopmentCards.BottomStack.PutNewCardToEmptySite(card)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New(fmt.Sprintf("你这个发展卡等级好像有问题，等级 = %d", cardLevel))
}

// ShowInfo 展示信息
func (t *Table) ShowInfo() {
	fmt.Printf("|=========Table=========\n")
	fmt.Printf("| Token: %+v\n", t.TokenStack)
	fmt.Printf("| DevCard: %+v\n", t.DevelopmentCardStacks.ShowIdxInfo())
	fmt.Printf("| RevealCards: %+v\n", t.RevealedDevelopmentCards.ShowIdxInfo())
	fmt.Printf("| Noble: %+v\n", t.NobleTilesStack.ShowIdxInfo())
	fmt.Printf("|=======================\n")
}
