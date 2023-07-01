package conponent

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

type RoomList struct {
	Id           int    `json:"id"`             // 房间序号
	RoomName     string `json:"room_name"`      // 房间名字
	CurPlayerNum int    `json:"cur_player_num"` // 当前玩家人数
	MaxPlayerNum int    `json:"max_player_num"` // 房间最大人数
}

func (r *RoomList) ToContainer() *fyne.Container {
	return container.New(layout.NewHBoxLayout(),
		widget.NewLabel(strconv.Itoa(r.Id)),
		widget.NewLabel(r.RoomName),
		widget.NewLabel(fmt.Sprintf("%d / %d", r.CurPlayerNum, r.MaxPlayerNum)),
		layout.NewSpacer(),
		widget.NewButton("Join", func() {
		}),
	)
}

func (r *RoomList) ToRowContainer() *fyne.Container {
	return container.NewGridWithColumns(5,
		widget.NewLabel(strconv.Itoa(r.Id)),
		widget.NewLabel(r.RoomName),
		widget.NewLabel(fmt.Sprintf("%d / %d", r.CurPlayerNum, r.MaxPlayerNum)),
		layout.NewSpacer(),
		widget.NewButton("Join", func() {}),
	)
}

func GetTestRoomList() []*RoomList {
	return []*RoomList{
		&RoomList{
			Id:           1,
			RoomName:     "test1",
			CurPlayerNum: 1,
			MaxPlayerNum: 4,
		},
		&RoomList{
			Id:           2,
			RoomName:     "test2",
			CurPlayerNum: 3,
			MaxPlayerNum: 4,
		},
	}
}
