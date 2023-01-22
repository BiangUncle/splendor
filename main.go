package main

import (
	"splendor/model"
	"splendor/service"
)

func init() {
	model.CsvFilePath = "csv/"
	model.Init()
}

func NewGame() {
	table, err := service.CreateANewGame(2)
	if err != nil {
		panic(err)
	}
	err = service.TurnRound(table)
	if err != nil {
		panic(err)
	}
}

func main() {
	NewGame()
}
