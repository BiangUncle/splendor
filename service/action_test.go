package service

import (
	"splendor/model"
	"testing"
)

func TestTurnAction(t *testing.T) {
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
