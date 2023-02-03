package model

import (
	"fmt"
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
	fmt.Println(table.ShowTableInfo())
}

func TestTable_Status(t *testing.T) {
	table := CreateTable()
	player := CreatePlayer()
	table.AddPlayer(player)
	table.Reveal()

	s, err := table.Status()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)

	newTable := CreateTable()
	err = newTable.LoadStatus(s)
	if err != nil {
		t.Fatal(err)
	}

	table.ShowVisualInfo()
	newTable.ShowVisualInfo()
}
