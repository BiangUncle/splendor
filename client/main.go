package main

import (
	"fmt"
	"os"
	"os/exec"
	"splendor/utils"
)

func Clf() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Start() {

	InitDefaultAction()

	c := ConstructClient()
	g := ConstructGameStatus(c)

	for {
		//Clf()
		g.Append(g.ShowPlayerInfo(), g.ReturnContent(), ShowOptionsInfos())
		g.Print()
		g.Clf()

		action, err := utils.InputInt("请选择你的操作")
		if err != nil {
			fmt.Println(err)
			continue
		}
		LoopAction(action)(g)
		if action == 0 {
			fmt.Println("exit")
			break
		}
	}

	return
}

func main() {
	Start()
}
