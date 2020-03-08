package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type TunnelServer struct {
	ctx         context.Context
	lockTunnels sync.Mutex
	tunnels     map[string]*websocket.Conn
}

func (s *TunnelServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id := request.Header.Get("ID")
	con, err := Upgrader.Upgrade(writer, request, nil)
	if err != nil {
		return
	}
	s.AddTunnelCon(con, id)
}

func NewTunnelServer(ctx context.Context) *TunnelServer {
	return &TunnelServer{
		lockTunnels: sync.Mutex{},
		tunnels:     make(map[string]*websocket.Conn, 100),
		ctx:         ctx,
	}
}

func (s *TunnelServer) AddTunnelCon(con *websocket.Conn, id string) {
	s.lockTunnels.Lock()
	fmt.Printf("Get a new connect connection from %v\n", id)
	s.tunnels[id] = con
	s.lockTunnels.Unlock()

	handleConnection(s.ctx, con)
}

func (s *TunnelServer) GetTunnelCon(id string) (*websocket.Conn, bool) {
	s.lockTunnels.Lock()
	defer s.lockTunnels.Unlock()
	con, ok := s.tunnels[id]
	return con, ok
}

func handleConnection(ctx context.Context, con *websocket.Conn) {
	defer con.Close()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("read handleSingleConnection loop stop")
			return
		default:
		}
		msgType, data, err := con.ReadMessage()
		if err != nil {
			fmt.Printf("read msg type %v error: %v\n", msgType, err)
			return
		}
		fmt.Printf("from %v get message type %v data:%v\n", con.RemoteAddr(), msgType, string(data))
	}
}
