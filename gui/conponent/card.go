package conponent

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"splendor/model"
)

func CardModel2UI(f fyne.Window, c *model.DevelopmentCard) *fyne.Container {
	cardM := widget.NewCard(
		fmt.Sprintf("%d", c.Idx),
		fmt.Sprintf("%d", c.Prestige),
		widget.NewButton("select", func() {
			fmt.Printf("你选择了卡 %+v\n", c.Idx)
		}),
	)
	cardM.Resize(fyne.NewSize(300, 300))
	return container.NewCenter(cardM)
}

func NobleModel2UI(f fyne.Window, n *model.NobleTile) *fyne.Container {

	nobleM := widget.NewCard(
		fmt.Sprintf("%d", n.Idx),
		fmt.Sprintf("%d", n.Prestige),
		widget.NewButton("select", func() {
			fmt.Printf("你选择了贵族 %+v\n", n.Idx)
		}),
	)
	nobleM.Resize(fyne.NewSize(300, 300))

	return container.NewCenter(nobleM)
}

//func CardButton(f fyne.Window) *widget.Button {
//	//btn := widget.NewButton("select", func() {
//	//	base.NewConfirmWindowWithContent(f, content).Show()
//	//})
//	file, err := os.Open("./gui/static/pics/card.png")
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//
//	var b []byte
//	_, err = file.Read(b)
//	if err != nil {
//		panic(err)
//	}
//	pics := fyne.NewStaticResource("pics", b)
//	//image := canvas.NewImageFromImage(img)
//	//c := container.NewCenter(img)
//	btn2 := widget.NewButton("", func() {
//
//	})
//
//	btn := widget.NewButtonWithIcon("", pics, func() {
//
//	})
//	return btn
//}
