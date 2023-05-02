package main

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"splendor/utils"
	"time"
)

var stopChannel = make(chan string, 0)

//Client custom business
func pingLoop(conn ziface.IConnection) {
	for {
		err := conn.SendMsg(1, []byte("Ping...Ping...Ping...[FromClient]"))
		if err != nil {
			fmt.Println(err)
			break
		}

		time.Sleep(1 * time.Second)
	}
}

func action(conn ziface.IConnection) {
	for {
		s, _ := utils.InputString("请输入操作")
		switch s {
		case "server":
			fmt.Println("加入服务器")
			err := conn.SendMsg(1, []byte("加入服务器"))
			if err != nil {
				fmt.Println(err)
				break
			}
		case "table":
			fmt.Println("加入房间")
			err := conn.SendMsg(1, []byte("加入房间"))
			if err != nil {
				fmt.Println(err)
				break
			}
		case "stop":
			fmt.Println("退出")
			stopChannel <- "stop"
			break
		}
	}
}

//Executed when a connection is created
func onClientStart(conn ziface.IConnection) {
	fmt.Println("onClientStart is Called ... ")
	go action(conn)
}

func main() {
	//Create a client client
	client := znet.NewClient("127.0.0.1", 7777)

	//Set the hook function after the link is successfully established
	client.SetOnConnStart(onClientStart)

	//start the client
	client.Start()

	<-stopChannel
	close(stopChannel)
}
