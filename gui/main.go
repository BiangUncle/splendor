package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"splendor/gui/page"
)

func InitWindow(w fyne.Window) {
	w.SetTitle("Splendor")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(1280, 720))
}

func main() {
	a := app.New()
	w := a.NewWindow("window")
	InitWindow(w)

	//roomListPage := page.NewRoomListPage()
	// 加载桌面
	gamePage := page.NewGamePage()
	w.SetContent(gamePage.Content(w))

	//confirmWindow := a.NewWindow("Confirm")
	//confirmWindow.SetContent(container.NewVBox(
	//	widget.NewLabel("Are you sure?"),
	//	container.NewHBox(
	//		widget.NewButton("yes", func() {
	//			fmt.Println("ok")
	//			confirmWindow.Close()
	//		}),
	//		widget.NewButton("no", func() {
	//			fmt.Println("no")
	//			confirmWindow.Close()
	//		}),
	//	),
	//))
	//menu1 := fyne.NewMenu("file",
	//	fyne.NewMenuItem("open", func() {
	//		fmt.Println("open")
	//	}), fyne.NewMenuItem("save", func() {
	//		fmt.Println("save")
	//	}),
	//)
	//popupmenu1 := widget.NewPopUpMenu(menu1, w.Canvas()) // 弹出菜单
	//btn1 := widget.NewButton("click", func() {
	//	popupmenu1.Show() // 显示
	//})
	//btn2 := widget.NewButton("click2", func() {
	//	popupmenu1.ShowAtPosition(fyne.NewPos(60, 60)) // 显示定位地方
	//})

	w.ShowAndRun()
}
