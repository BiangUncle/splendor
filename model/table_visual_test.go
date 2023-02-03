package model

import (
	"testing"
)

func TestTable_ShowVisualInfo(t1 *testing.T) {
	table := CreateTable()
	player := CreatePlayer()
	table.AddPlayer(player)
	table.Reveal()

	table.ShowVisualInfo()
}

func TestTable_Marshal(t *testing.T) {
	table := CreateTable()
	player := CreatePlayer()
	table.AddPlayer(player)
	table.Reveal()

	s, err := table.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)

	newTable := CreateTable()
	err = newTable.Unmarshal(s)
	if err != nil {
		t.Fatal(err)
	}

	table.ShowVisualInfo()
	newTable.ShowVisualInfo()
}
