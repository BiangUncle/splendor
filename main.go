package main

import (
	"fmt"
	"splendor/model"
	"splendor/service"
)

func init() {
	model.CsvFilePath = "csv/"
	model.Init()
	model.InitDefaultTable()
}

func RunServer() {
	service.Run()
}

func TableVisual() {
	table := model.CreateTable()
	player := model.CreatePlayer()
	table.AddPlayer(player)
	table.Reveal()

	table.ShowVisualInfo()
}

func Whole() {
	ds := model.CreateTable()
	ds.Reveal()

	for _, s := range ds.WholeVisual() {
		fmt.Println(s)
	}
}

func main() {
	//NewGame()
	//RunServer()
	//TableVisual()
	Whole()
}
