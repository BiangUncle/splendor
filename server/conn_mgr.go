package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

var onceConnManager = &sync.Once{}
var connManager *ConnManager

type ConnManager struct {
	conn map[int]*websocket.Conn
	id   int
	l    *sync.RWMutex
}

func NewConnManager() *ConnManager {
	onceConnManager.Do(func() {
		connManager = &ConnManager{
			conn: make(map[int]*websocket.Conn),
			id:   1,
			l:    &sync.RWMutex{},
		}
	})
	return connManager
}

func (m *ConnManager) AddConn(conn *websocket.Conn) {
	m.l.Lock()
	defer m.l.Unlock()
	m.conn[m.id] = conn
	m.id++
}

func (m *ConnManager) Close() {
	for id, conn := range m.conn {
		conn.Close()
		fmt.Printf("关闭连接[%d], conn = %+v\n", id, conn)
	}
}
