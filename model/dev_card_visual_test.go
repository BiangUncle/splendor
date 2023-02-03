package model

import (
	"fmt"
	"testing"
)

func TestDevelopmentCard_WholeCard(t *testing.T) {
	ds := CreateANewDevelopmentCardStacks()
	c, err := ds.TopStack.TakeTopCard()
	if err != nil {
		t.Fatal(err)
	}

	for _, s := range c.WholeCard() {
		fmt.Println(s)
	}
}
