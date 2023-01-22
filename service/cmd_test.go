package service

import (
	"testing"
)

func TestTurnRoundCMD(t *testing.T) {
	table, err := CreateANewGame(2)
	if err != nil {
		t.Fatal(err)
	}
	err = TurnRound(table)
	if err != nil {
		t.Fatal(err)
	}
}
