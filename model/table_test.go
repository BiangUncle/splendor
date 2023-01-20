package model

import (
	"testing"
)

func TestCreateANewTable(t *testing.T) {
	table := CreateANewTable()
	t.Logf("%+v", table)
}
