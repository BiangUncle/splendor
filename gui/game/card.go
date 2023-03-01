package game

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Card struct {
	app.Compo
	Name string
}

func (c *Card) Render() app.UI {
	return app.Div().
		Class("dev_card").OnClick(c.onGameSelect).Body(app.Text(c.Name))
}

func (c *Card) onGameSelect(ctx app.Context, e app.Event) {
	fmt.Println("select", c.Name)
}
