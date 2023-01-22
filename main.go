package main

import (
	"splendor/model"
	"splendor/service"
)

func init() {
	model.CsvFilePath = "csv/"
	model.Init()
}

func main() {
	table, err := service.CreateANewGame(2)
	if err != nil {
		panic(err)
	}
	err = service.TurnRoundCMD(table)
	if err != nil {
		panic(err)
	}
}
