package service

import (
	"splendor/model"
	"testing"
)

func TestMain(m *testing.M) {
	model.Init()
	m.Run()
}
