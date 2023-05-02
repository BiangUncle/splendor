package main

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"splendor/model"
	"splendor/utils"
)

type PingRouter struct {
	znet.BaseRouter
}

type GetTableListRouter struct {
	znet.BaseRouter
}

func (r *GetTableListRouter) Handle(request ziface.IRequest) {

}

//Ping Handle MsgId=1
func (r *PingRouter) Handle(request ziface.IRequest) {
	//read client data
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
}

// Input 服务器接受输入
func Input() {
	for {
		s, _ := utils.InputString("")
		if s == "table" {
			fmt.Println(model.GetDefaultTable().ShowTableInfo())
			continue
		}
		break
	}
}

func main() {
	// 常见服务器对象
	s := znet.NewServer()

	// 配置路由
	s.AddRouter(1, &PingRouter{})
	s.AddRouter(2, &GetTableListRouter{})

	// 启动服务
	go s.Serve()

	Input()

	s.Stop()
}
