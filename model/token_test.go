package model

import (
	"fmt"
	"github.com/fatih/color"
	"testing"
)

func TestTokenStack_Color(t *testing.T) {
	/*
		for b := 40; b <= 47; b++ { // 背景色彩 = 40-47
			fmt.Printf("\033[%dm \033[0m=%d ", b, b)
		}
		fmt.Println()
	*/
	token := CreatANewTokenStack()
	c := color.New()

	for i, v := range token {
		c.Add(ColorConfig[i])
		fmt.Printf("%sx%d\n", c.Sprintf(" "), v)
	}
}
