package main

import (
	"fmt"
	"splendor/utils"
)

func inputString() string {
	s, err := utils.InputString()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return s
}

func inputInt() int {
	i, err := utils.InputInt()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return i
}
