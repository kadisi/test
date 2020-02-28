package main

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

type SessionManager struct {
	lockTunnels sync.Mutex
	tunnels     map[string]*websocket.Conn
	ctx         context.Context
}

func (s *SessionManager) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	con, err := Upgrader.Upgrade(writer, request, nil)
	if err != nil {
		return
	}

	s.AddTunnelCon(con)
}

func NewsessionManager(ctx context.Context) *SessionManager {
	return &SessionManager{
		lockTunnels: sync.Mutex{},
		tunnels:     make(map[string]*websocket.Conn, 100),
		ctx:         ctx,
	}
}

func (s *SessionManager) AddTunnelCon(con *websocket.Conn) {
	remote := con.RemoteAddr().String()
	s.lockTunnels.Lock()
	s.tunnels[remote] = con
	fmt.Printf("Get a new connect connection %v\n", remote)
	//go handleConnection(s.ctx, con)
	s.lockTunnels.Unlock()
}

func handleConnection(ctx context.Context, con *websocket.Conn) {
	defer con.Close()
	for {
		select {
		case <-ctx.Done():
			con.WriteControl(websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, "i will be closed"), time.Now().Add(time.Second))
			fmt.Println("handleSingleConnection loop stop")
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
