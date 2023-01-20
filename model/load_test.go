package model

import (
	"testing"
)

func TestLoad(t *testing.T) {
	LoadDefaultDevelopmentCard()
	LoadDefaultNobleTiles()

	c := defaultDevelopmentCardStacks
	n := defaultNobleTilesStack
	_ = c
	_ = n
	t.Log("success")
}
