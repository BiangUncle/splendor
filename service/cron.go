package service

import (
	"fmt"
	"splendor/model"
	"sync"
	"time"
)

type OneCron struct {
	stop chan struct{}
	wg   *sync.WaitGroup
}

func ConstructCron() *OneCron {
	return &OneCron{
		stop: make(chan struct{}),
		wg:   &sync.WaitGroup{},
	}
}

func (o *OneCron) CheckPlayerNetStatus() {
	fmt.Println("检查协程: ", "检查桌台，用户状态")
	model.ClearDefaultTable()
	model.ClearOfflinePlayer()
}

func (o *OneCron) Run() {
	t1 := time.Tick(10 * time.Second)
	for {
		select {
		case <-t1:
			o.CheckPlayerNetStatus()
		case <-o.stop:
			fmt.Println("心跳协程: ", "关闭保活协程")
			close(o.stop)
			return
		}
	}
}

func (o *OneCron) Stop() {
	o.stop <- struct{}{}
}
