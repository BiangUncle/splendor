package service

import (
	"splendor/model"
	"testing"
)

func TestPlayerAddToken(t *testing.T) {
	table := model.CreateANewTable()
	t.Logf("%+v", table)

	player1 := model.CreateANewPlayer()
	player1.ShowPlayerInfo()

	err := TurnAction(player1, table, 1)
	if err != nil {
		t.Error(err)
	}
	player1.ShowPlayerInfo()
	t.Logf("%+v", table)
}

func TestShuffleDevCard(t *testing.T) {
	table := model.CreateANewTable()
	t.Logf("%+v", table)

	table.DevelopmentCardStacks.ShowIdxInfo()
	table.Shuffle()
	table.DevelopmentCardStacks.ShowIdxInfo()
}

func TestShuffleNobleTile(t *testing.T) {
	table := model.CreateANewTable()
	t.Logf("%+v", table)

	table.NobleTilesStack.ShowIdxInfo()
	table.Shuffle()
	table.NobleTilesStack.ShowIdxInfo()
}

func TestTakeTopCard(t *testing.T) {
	table := model.CreateANewTable()
	t.Logf("%+v", table)

	err := table.Reveal()
	if err != nil {
		t.Error(err)
	}

	err = table.Reveal()
	if err != nil {
		t.Error(err)
	}
}

func TestPurchaseDevelopmentCard(t *testing.T) {
	table := model.CreateANewTable()
	t.Logf("%+v\n", table.RevealedDevelopmentCards.ShowIdxInfo())
	t.Logf("%+v\n", table.DevelopmentCardStacks.ShowIdxInfo())

	err := table.Reveal()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", table.RevealedDevelopmentCards.ShowIdxInfo())
	t.Logf("%+v\n", table.DevelopmentCardStacks.ShowIdxInfo())

	player1 := model.CreateANewPlayer()
	player1.ShowPlayerInfo()

	err = PurchaseDevelopmentCard(player1, table, 10001) // todo 可配置
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", table.RevealedDevelopmentCards.ShowIdxInfo())
	t.Logf("%+v\n", table.DevelopmentCardStacks.ShowIdxInfo())

	player1.ShowPlayerInfo()
}

func TestPurchaseDevelopmentCardWithToken(t *testing.T) {
	table := model.CreateANewTable()
	table.ShowInfo()

	err := table.Reveal()
	if err != nil {
		t.Fatal(err)
	}

	table.ShowInfo()

	player1 := model.CreateANewPlayer()
	player1.Tokens.Add([]int{3, 3, 3, 3, 3, 0})
	player1.ShowPlayerInfo()

	err = PurchaseDevelopmentCard(player1, table, 10001) // todo 可配置
	if err != nil {
		t.Fatal(err)
	}

	table.ShowInfo()

	player1.ShowPlayerInfo()
}
