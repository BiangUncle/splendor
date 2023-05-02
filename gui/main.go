package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("window")
	w.SetTitle("Splendor")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(1280, 720))
	txtLabel := widget.NewLabel("Room List")
	butCreateRoom := widget.NewButton("Create Room", func() {})
	title := container.New(layout.NewHBoxLayout(), txtLabel, layout.NewSpacer(), butCreateRoom)
	w.SetContent(container.NewVBox(
		title,
	))

	w.ShowAndRun()
}
