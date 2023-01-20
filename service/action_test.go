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
