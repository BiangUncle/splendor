package page

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"splendor/gui/conponent"
	"splendor/model"
)

type GamePage struct {
	CardStack  *CardStackComponent
	NobleStack *NobleStackComponent
}

// CardStackComponent 卡堆组件
type CardStackComponent struct {
}

// NobleStackComponent 贵族组件
type NobleStackComponent struct {
}

// TokenComponent 宝石组件
type TokenComponent struct {
}

func NewGamePage() *GamePage {
	return new(GamePage)
}

func (p *GamePage) GameMenus() *fyne.Container {
	return container.New(layout.NewHBoxLayout(),
		widget.NewLabel("Room Name"),
		layout.NewSpacer(),
		widget.NewLabel("00:00:12"),
		//widget.NewButton("Refresh", func() {}),
	)
}

func (p *GamePage) Content(f fyne.Window) *fyne.Container {
	return container.NewVBox(
		p.GameMenus(),
		p.NobleStack.Content(f),
		p.CardStack.Content(f),
	)
}

func (c *CardStackComponent) Content(f fyne.Window) *fyne.Container {
	hb := container.NewVBox(
		container.NewHBox(
			conponent.CardModel2UI(f, &model.DevelopmentCard{
				Idx: 10001,
			}),
			conponent.CardModel2UI(f, &model.DevelopmentCard{
				Idx: 10002,
			}),
			conponent.CardModel2UI(f, &model.DevelopmentCard{
				Idx: 10003,
			}),
			conponent.CardModel2UI(f, &model.DevelopmentCard{
				Idx: 10004,
			}),
		),
		container.NewHBox(
			conponent.CardModel2UI(f, &model.DevelopmentCard{
				Idx: 10005,
			}),
			conponent.CardModel2UI(f, &model.DevelopmentCard{
				Idx: 10006,
			}),
			conponent.CardModel2UI(f, &model.DevelopmentCard{
				Idx: 10007,
			}),
			conponent.CardModel2UI(f, &model.DevelopmentCard{
				Idx: 10008,
			}),
		),
		container.NewHBox(
			conponent.CardModel2UI(f, &model.DevelopmentCard{
				Idx: 10009,
			}),
			conponent.CardModel2UI(f, &model.DevelopmentCard{
				Idx: 10010,
			}),
			conponent.CardModel2UI(f, &model.DevelopmentCard{
				Idx: 10011,
			}),
			conponent.CardModel2UI(f, &model.DevelopmentCard{
				Idx: 10012,
			}),
		),
	)

	return hb
}

func (n *NobleStackComponent) Content(f fyne.Window) *fyne.Container {
	hb := container.NewVBox(
		container.NewHBox(
			conponent.NobleModel2UI(f, &model.NobleTile{
				Idx: 20001,
			}),
			conponent.NobleModel2UI(f, &model.NobleTile{
				Idx: 20002,
			}),
			conponent.NobleModel2UI(f, &model.NobleTile{
				Idx: 20003,
			}),
		),
	)
	return hb
}

func (t *TokenComponent) Content() *fyne.Container {
	//t := container.NewVBox(
	//	widget.NewButton("7", func() {}),
	//)
	//return circle
	return nil
}
