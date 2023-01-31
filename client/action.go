package main

import (
	"errors"
	"fmt"
)

/*
MIND: 将所有的Action出现的报错打印集成在action中，不在外面进行打印
*/

type Action struct {
	gs       *GameStatus
	tryTime  int
	exitFunc func(interface{}) bool
}

var InvalidExitError = errors.New("手动退出循环错误")

// ReturnTokens 返还宝石循环操作
func (a *Action) ReturnTokens() error {
	for {
		a.tryTime++
		fmt.Print("请输入要丢弃的宝石: ")
		tokensString := inputString()

		if a.CheckExit(tokensString) {
			return InvalidExitError
		}
		_, err := a.gs.ReturnTokens(tokensString)
		if err != nil {
			fmt.Println(err)
			continue
		}
		return nil
	}
}

func (a *Action) SetExitFunc(f func(interface{}) bool) *Action {
	a.exitFunc = f
	return a
}

func (a *Action) CheckExit(signal interface{}) bool {
	if a.exitFunc == nil {
		return false
	}
	return a.exitFunc(signal)
}
