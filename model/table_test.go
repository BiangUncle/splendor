package model

import (
	"testing"
)

func TestCreateANewTable(t *testing.T) {
	table := CreateTable()
	t.Logf("%+v", table)
}

func TestTable_ShowTableInfo(t *testing.T) {
	table := CreateTable()
	player1 := CreatePlayer()
	player2 := CreatePlayer()
	table.AddPlayer(player1)
	table.AddPlayer(player2)

	err := table.Reveal()
	if err != nil {
		t.Fatal(err)
	}
	table.ShowTableInfo()
}
