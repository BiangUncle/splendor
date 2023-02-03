package model

import (
	"fmt"
	"github.com/fatih/color"
	"strconv"
)

// ShowIdxInfo 展示牌的索引
func (s DevelopmentCardStack) ShowIdxInfo() string {
	idxInfo := make([]int, len(s))
	for i, card := range s {
		if card == nil {
			idxInfo[i] = -1
		} else {
			idxInfo[i] = card.Idx % 100
		}
	}
	return fmt.Sprintf("%+v", idxInfo)
}

// ShowIdxInfo 展示牌的索引
func (s *DevelopmentCardStacks) ShowIdxInfo() string {
	return s.TopStack.ShowIdxInfo() + s.MiddleStack.ShowIdxInfo() + s.BottomStack.ShowIdxInfo()
}

func (c *DevelopmentCard) Visual() string {
	require := ""
	p := color.New()
	typeCount := 0

	for idx, v := range c.Acquires {
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
	bonusC := color.New(ColorConfig[c.BonusType])

	return fmt.Sprintf("[%s  %-4s]", bonusC.Sprintf("%d", c.Prestige), require)
}

func (c *DevelopmentCard) VisualV2() string {
	require := ""
	p := color.New(ColorConfig[c.BonusType])

	for _, v := range c.Acquires {
		require += strconv.Itoa(v)
	}

	require = p.Sprintf("%d  %s", c.Prestige, require)

	return require
}

func (s DevelopmentCardStack) Visual() string {
	info := ""
	for _, c := range s {
		if c == nil {
			info += "[       ]"
			continue
		}
		info += c.Visual()
	}
	return info
}

func (s DevelopmentCardStacks) Visual() string {
	info := ""
	info += s.TopStack.Visual() + "\n"
	info += s.MiddleStack.Visual() + "\n"
	info += s.BottomStack.Visual()
	return info
}

/*
+-----+
|12  2|
|     |
|     |
|02111|
+-----+
*/

func (c *DevelopmentCard) WholeCard() []string {

	bonusC := color.New(ColorConfig[c.BonusType])

	require := ""
	p := color.New()
	typeCount := 0

	for idx, v := range c.Acquires {
		if idx == 5 {
			continue
		}
		typeCount++
		p.Add(ColorConfig[idx])
		require += p.Sprintf("%d", v)
	}

	p = color.New(ColorConfig[c.BonusType])

	var ret []string
	ret = append(ret, "+-----+")
	ret = append(ret, fmt.Sprintf("|%s|", p.Sprintf("%-2d  %s", c.Idx%10000, bonusC.Sprintf("%d", c.Prestige))))
	ret = append(ret, fmt.Sprintf("|%s|", p.Sprint("     ")))
	ret = append(ret, fmt.Sprintf("|%s|", p.Sprint("     ")))
	ret = append(ret, fmt.Sprintf("|%s|", require))
	ret = append(ret, "+-----+")

	return ret
}

func (s DevelopmentCardStack) WholeCard() []string {
	ret := make([]string, 6)
	for i := 0; i < 6; i++ {
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
