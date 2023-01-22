package model

import (
	"testing"
)

func TestCreateANewPlayer(t *testing.T) {
	player := CreateANewPlayer()
	player.ShowPlayerInfoV2()
}
