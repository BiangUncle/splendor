package game

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Game struct {
	app.Compo
}

func (g *Game) Render() app.UI {
	return GetTable("1")
}
