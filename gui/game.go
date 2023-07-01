package main

import (
	"fyne.io/fyne/v2"
	"splendor/gui/net"
)

type Game struct {
	Client *net.Client
	Window fyne.Window
}
