package utils

import (
	"fmt"
	"github.com/fatih/color"
)

func _systemPrefix() {
	c := color.New(color.FgCyan)
	fmt.Printf("[%s]", c.Sprint("SYSTEM"))
}

func SystemPrint(a ...interface{}) {
	_systemPrefix()
	fmt.Println(a...)
}

func SystemPrintf(format string, a ...any) {
	_systemPrefix()
	fmt.Printf(format, a...)
}
