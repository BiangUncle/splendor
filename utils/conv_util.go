package utils

import "strconv"

// ToInt 字符串转数字
func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return i
}
