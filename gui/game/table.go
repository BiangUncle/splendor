package game

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Table struct {
	app.Compo

	TopCardStack    []*Card
	MiddleCardStack []*Card
	BottomCardStack []*Card
}

var OnlineTable *Table

func init() {
	OnlineTable = CreateNewTable()
}

func GetTable(tableId string) *Table {
	return OnlineTable
}

func CreateNewTable() *Table {
	tb := &Table{}
	tb.TopCardStack = []*Card{
		&Card{Name: "Card1"},
		&Card{Name: "Card2"},
		&Card{Name: "Card3"},
		&Card{Name: "Card4"},
	}
	fmt.Println("init top card stack")
	tb.MiddleCardStack = []*Card{
		&Card{Name: "Card1"},
		&Card{Name: "Card2"},
		&Card{Name: "Card3"},
		&Card{Name: "Card4"},
	}
	fmt.Println("init middle card stack")
	tb.BottomCardStack = []*Card{
		&Card{Name: "Card1"},
		&Card{Name: "Card2"},
		&Card{Name: "Card3"},
		&Card{Name: "Card4"},
	}
	fmt.Println("init bottom card stack")
	return tb
}

func (tb *Table) Render() app.UI {
	page := app.Div().Class("table").Body(
		app.Div().Class("card_stack top_card_stack").Body(
			app.Range(tb.TopCardStack).Slice(func(i int) app.UI {
				return tb.TopCardStack[i]
			}),
		),
		app.Div().Class("card_stack middle_card_stack").Body(
			app.Range(tb.MiddleCardStack).Slice(func(i int) app.UI {
				return tb.MiddleCardStack[i]
			}),
		),
		app.Div().Class("card_stack bottom_card_stack").Body(
			app.Range(tb.BottomCardStack).Slice(func(i int) app.UI {
				return tb.BottomCardStack[i]
			}),
		),
	)
	return page
}
