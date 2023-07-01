package base

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ConfirmWidget interface {
	PopWidget
}

func ConfirmWindow(w fyne.Window) *dialog.ConfirmDialog {
	confirm := dialog.NewConfirm("Are you sure", "", func(b bool) {
		if b {
			fmt.Println("ok")
		} else {
			fmt.Println("no")
		}
	}, w)
	return confirm
}

type ConfirmWindowWithContent struct {
	pop *widget.PopUp
}

func NewConfirmWindowWithContent(f fyne.Window, innerContent fyne.CanvasObject) *ConfirmWindowWithContent {

	c := &ConfirmWindowWithContent{}

	dismiss := &widget.Button{Text: "No", Icon: theme.CancelIcon(),
		OnTapped: func() {
			c.Hide()
		},
	}

	confirm := &widget.Button{Text: "Yes", Icon: theme.ConfirmIcon(), Importance: widget.HighImportance,
		OnTapped: func() {
			c.Hide()
		},
	}

	content := container.NewVBox(
		innerContent,
		container.NewHBox(confirm, dismiss),
	)

	pop := widget.NewModalPopUp(content, f.Canvas())
	pop.Refresh()
	c.pop = pop

	return c
}

func (c *ConfirmWindowWithContent) Show() {
	c.pop.Show()
}

func (c *ConfirmWindowWithContent) Hide() {
	c.pop.Hide()
}
