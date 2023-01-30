package main

import (
	"fmt"
	"time"
)

type GameCron struct {
	stop chan struct{}
}

func (g *GameCron) Stop() {
	g.stop <- struct{}{}
}

func (g *GameStatus) RoutineKeepALive() {
	t1 := time.Tick(3 * time.Second)
	for {
		select {
		case <-t1:
			_, err := g.KeepALive()
			if err != nil {
				fmt.Println("心跳协程: ", "保活出现问题")
			}
		case <-g.stop:
			fmt.Println("心跳协程: ", "关闭保活协程")
			close(g.stop)
			return
		}
	}
}
