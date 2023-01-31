package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputList() ([]int, error) {
	var num []int
	reader := bufio.NewReader(os.Stdin)
	strBytes, _, err := reader.ReadLine()
	if err != nil {
		return nil, err
	}
	str := strings.Fields(string(strBytes))
	for i := range str {
		n, err := strconv.Atoi(str[i])
		if err != nil {
			return nil, err
		}
		num = append(num, n)
	}
	return num, nil
}

func InputInt() (int, error) {
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func InputString() (string, error) {
	var s string
	_, err := fmt.Scanf("%s", &s)
	if err != nil {
		return "", err
	}
	return s, nil
}
