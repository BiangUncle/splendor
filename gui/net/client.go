package net

import (
	"bufio"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"os"
	pb "splendor/pb/proto"
	"sync"
)

var once = &sync.Once{}
var client *Client

type Client struct {
	conn *websocket.Conn
}

func NewClient() *Client {
	once.Do(func() {
		client = &Client{}
	})
	return client
}

func (c *Client) AddConn(conn *websocket.Conn) {
	c.conn = conn
}

func (c *Client) Stop() {
	c.conn.Close()
}

func (c *Client) Conn() {
	dl := websocket.Dialer{}
	conn, _, err := dl.Dial("ws://127.0.0.1:8888", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		m, p, e := conn.ReadMessage()
		if e != nil {
			break
		}
		fmt.Println("get message type = ", m)
		if m == 1 {
			roomList := &pb.RoomList{}
			err = proto.Unmarshal(p, roomList)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(roomList)
		}
	}
}

func (c *Client) Run() {
	for {
		reader := bufio.NewReader(os.Stdin)
		l, _, _ := reader.ReadLine()
		c.conn.WriteMessage(1, l)
	}
}

func (c *Client) SendMsg(messageType int, msg []byte) {
	err := c.conn.WriteMessage(messageType, msg)
	if err != nil {
		fmt.Println("send message error, ", err)
	}
}
