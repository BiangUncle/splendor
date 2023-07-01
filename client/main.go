package main

import (
	"bufio"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"os"
	pb "splendor/pb/proto"
)

func main() {
	dl := websocket.Dialer{}
	conn, _, err := dl.Dial("ws://127.0.0.1:8888", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	go send(conn)

	for {
		_, p, e := conn.ReadMessage()
		if e != nil {
			break
		}
		roomList := &pb.RoomList{}
		err = proto.Unmarshal(p, roomList)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(roomList)
	}
}

func send(conn *websocket.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		l, _, _ := reader.ReadLine()
		conn.WriteMessage(1, l)
	}
}
