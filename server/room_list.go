package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "splendor/pb/proto"
)

func GetRoomList() []byte {
	roomList := &pb.RoomList{
		RoomInfo: []*pb.RoomInfo{
			&pb.RoomInfo{
				RoomId:       1,
				RoomName:     "test1",
				CurPlayerNum: 1,
				MaxPlayerNum: 4,
			},
			&pb.RoomInfo{
				RoomId:       2,
				RoomName:     "test2",
				CurPlayerNum: 3,
				MaxPlayerNum: 4,
			},
		},
	}

	data, err := proto.Marshal(roomList)
	if err != nil {
		fmt.Printf("error, err = %+v\n", err)
		return nil
	}
	return data
}
