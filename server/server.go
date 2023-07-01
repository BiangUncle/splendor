package main

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var s *Server
var once = &sync.Once{}
var DefaultUpGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Server struct {
	isShutDown  bool
	handler     *http.ServeMux
	server      *http.Server
	connManager *ConnManager
	up          websocket.Upgrader
	handleFunc  map[int]func([]byte, *websocket.Conn)
}

func NewServer() *Server {
	once.Do(func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", handler)
		server := &http.Server{
			Addr:    ":8888",
			Handler: mux,
		}
		s = &Server{
			isShutDown:  false,
			handler:     mux,
			server:      server,
			connManager: NewConnManager(),
			up:          DefaultUpGrader,
			handleFunc:  make(map[int]func([]byte, *websocket.Conn)),
		}
		s.Register(1, handleGetRoomList)
	})
	return s
}

func (s *Server) Run() {
	go s.server.ListenAndServe()
	listenSignal(context.Background(), s.server)
}

func handleGetRoomList(msg []byte, conn *websocket.Conn) {
	fmt.Println("handle msg =", string(msg))
	data := GetRoomList()
	conn.WriteMessage(1, data)
}

func handleMessage(conn *websocket.Conn) {
	for {
		m, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		// 判断消息支持注册
		if _, ok := s.handleFunc[m]; ok {
			s.handleFunc[m](p, conn)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := s.up.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	connManager.AddConn(conn)
	go handleMessage(conn)
}

// listenSignal 监听断开流程
func listenSignal(ctx context.Context, httpSrv *http.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-sigs:
		fmt.Println("notify sigs")
		s.Stop(ctx)
		fmt.Println("http shutdown")
	}
}

func (s *Server) Register(msgId int, f func([]byte, *websocket.Conn)) {
	s.handleFunc[msgId] = f
}

func (s *Server) Stop(ctx context.Context) {
	s.isShutDown = true
	s.connManager.Close()
	s.server.Close()
	s.server.Shutdown(ctx)
}
