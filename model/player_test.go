package model

import (
	"testing"
)

func TestCreateANewPlayer(t *testing.T) {
	player := CreatePlayer()
	player.ShowPlayerInfoV2()
}
