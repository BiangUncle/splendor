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

/*
□
*/

func (s TokenStack) WholeVisual() []string {
	var ret []string
	p := color.New()
	for idx, v := range s {
		t := ""
		count := 0
		for i := 0; i < v; i++ {
			t += "□"
			count++
		}
		p.Add(ColorConfig[idx])
		t = p.Sprint(t)
		for i := 0; i < 7-count; i++ {
			t += " "
		}
		ret = append(ret, fmt.Sprintf("[%s]", t))
	}
	return ret
}
