package page

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/golang/protobuf/proto"
	"splendor/gui/conponent"
	"splendor/gui/net"
	pb "splendor/pb/proto"
)

// RoomListPage 房间列表页面
type RoomListPage struct {
	Data   []*conponent.RoomList
	Client *net.Client
}

func NewRoomListPage(client *net.Client) *RoomListPage {
	return &RoomListPage{
		Data:   conponent.GetTestRoomList(),
		Client: client,
	}
}

func (p *RoomListPage) RoomListTitle() *fyne.Container {
	return container.New(layout.NewHBoxLayout(),
		widget.NewLabel("Room List"),
		layout.NewSpacer(),
		widget.NewButton("Create Room", func() {}),
		widget.NewButton("Refresh", func() {
			p.Client.GetRoomList()
		}),
	)
}

func (p *RoomListPage) NewRoomTable() *fyne.Container {
	l := container.NewGridWithRows(len(p.Data) + 1)
	l.Add(container.NewGridWithColumns(5,
		widget.NewLabel("Room Id"),
		widget.NewLabel("Room Name"),
		widget.NewLabel("Player"),
		layout.NewSpacer(),
		widget.NewLabel(""),
	))
	for _, room := range p.Data {
		l.Add(room.ToRowContainer())
	}
	return l
}

func (p *RoomListPage) Content() *fyne.Container {
	return container.NewVBox(
		p.RoomListTitle(),
		widget.NewSeparator(),
		p.NewRoomTable(),
	)
}

func (p *RoomListPage) UpdateRoomList(msg []byte) {
	roomList := &pb.RoomList{}
	err := proto.Unmarshal(msg, roomList)
	if err != nil {
		fmt.Println(err)
	}
	p.Data = make([]*conponent.RoomList, len(roomList.RoomInfo))
	for i, room := range roomList.RoomInfo {
		p.Data[i] = proto2model(room)
	}
}

func proto2model(p *pb.RoomInfo) *conponent.RoomList {
	return &conponent.RoomList{
		Id:           int(p.RoomId),
		RoomName:     p.RoomName,
		CurPlayerNum: int(p.CurPlayerNum),
		MaxPlayerNum: int(p.MaxPlayerNum),
	}
}
