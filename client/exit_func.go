package main

import "fmt"

type ExitFunc interface {
	IfExit(interface{}) bool
}

type StringCheckFunc struct{}
type IntCheckFunc struct{}

func (f *StringCheckFunc) IfExit(signal interface{}) bool {
	s := signal.(string)
	if s == "exit" {
		fmt.Println("手动退出循环")
		return true
	}
	return false
}

func (f *IntCheckFunc) IfExit(signal interface{}) bool {
	s := signal.(int)
	if s == -1 {
		fmt.Println("手动退出循环")
		return true
	}
	return false
}

func StringExitFunc(signal interface{}) bool {
	s := signal.(string)
	if s == "exit" {
		fmt.Println("手动退出循环")
		return true
	}
	return false
}

func IntExitFuncIfExit(signal interface{}) bool {
	s := signal.(int)
	if s == -1 {
		fmt.Println("手动退出循环")
		return true
	}
	return false
}
