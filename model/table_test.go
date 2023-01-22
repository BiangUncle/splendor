package model

import (
	"testing"
)

func TestCreateANewTable(t *testing.T) {
	table := CreateANewTable()
	t.Logf("%+v", table)
}

func TestTable_ShowTableInfo(t *testing.T) {
	table := CreateANewTable()
	player1 := CreateANewPlayer()
	player2 := CreateANewPlayer()
	table.AddPlayer(player1)
	table.AddPlayer(player2)

	err := table.Reveal()
	if err != nil {
		t.Fatal(err)
	}
	table.ShowTableInfo()
}
