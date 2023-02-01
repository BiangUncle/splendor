package model

import (
	"fmt"
	"github.com/fatih/color"
)

// ShowIdxInfo 展示信息
func (s NobleTilesStack) ShowIdxInfo() string {
	idxInfo := make([]int, len(s))
	for i, noble := range s {
		if noble == nil {
			idxInfo[i] = -1
		} else {
			idxInfo[i] = noble.Idx % 100
		}
	}
	//fmt.Printf("%+v\n", idxInfo)
	return fmt.Sprintf("%+v", idxInfo)
}

func (n *NobleTile) Visual() string {
	require := ""
	p := color.New()
	typeCount := 0

	for idx, v := range n.Acquires {
		if v == 0 {
			continue
		}
		typeCount++
		p.Add(ColorConfig[idx])
		if idx == TokenIdxOnyx {
			require += p.Sprintf("%s", color.WhiteString("%d", v))
		} else {
			require += p.Sprintf("%d", v)
		}
	}

	// 前面补充空格，保持一致
	for i := 0; i < (4 - typeCount); i++ {
		require = " " + require
	}

	return fmt.Sprintf("[%d  %-4s]", n.Prestige, require)
}

func (s NobleTilesStack) Visual() string {
	info := ""
	for _, n := range s {
		info += n.Visual()
	}
	return info
}
