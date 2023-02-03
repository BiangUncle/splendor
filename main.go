package main

import (
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

func main() {
	//NewGame()
	RunServer()
	//TableVisual()
}
