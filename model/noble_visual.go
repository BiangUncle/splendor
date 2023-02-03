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
		require += p.Sprintf("%d", v)
		//if idx == TokenIdxOnyx {
		//	require += p.Sprintf("%s", color.WhiteString("%d", v))
		//} else {
		//	require += p.Sprintf("%d", v)
		//}
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
		if n == nil {
			info += "[      ]"
		} else {
			info += n.Visual()
		}
	}
	return info
}

func (n *NobleTile) WholeCard() []string {

	require := ""
	p := color.New()
	typeCount := 0

	for idx, v := range n.Acquires {
		if idx == 5 {
			continue
		}
		typeCount++
		p.Add(ColorConfig[idx])
		require += p.Sprintf("%d", v)
	}

	var ret []string
	ret = append(ret, "+-----+")
	ret = append(ret, fmt.Sprintf("|%s|", p.Sprintf("%-2d  %d", n.Idx%10000, n.Prestige)))
	ret = append(ret, fmt.Sprintf("|%s|", require))
	ret = append(ret, "+-----+")

	return ret
}

func (s NobleTilesStack) WholeCard() []string {
	ret := make([]string, 4)
	for i := 0; i < 4; i++ {
		ret[i] = ""
	}

	for _, c := range s {
		wv := c.WholeCard()
		for idx, str := range wv {
			ret[idx] = ret[idx] + str
		}
	}
	return ret
}
