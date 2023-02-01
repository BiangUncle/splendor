package model

import (
	"fmt"
	"github.com/fatih/color"
)

func (s TokenStack) Visual() string {
	info := ""
	c := color.New()

	for i, v := range s {
		c.Add(ColorConfig[i])
		info += fmt.Sprintf("%sx%d ", c.Sprintf(" "), v)
	}

	return info
}
