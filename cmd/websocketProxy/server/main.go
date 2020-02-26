package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		w.WriteHeader(status)
		w.Write([]byte(reason.Error()))
	},
}

var ConMap map[string]*websocket.Conn

func init() {
	ConMap = make(map[string]*websocket.Conn, 10)
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		con, err := Upgrader.Upgrade(writer, request, nil)
		if err != nil {
			return
		}
		con.SetCloseHandler(func(code int, text string) error {
			fmt.Printf("detected closed %d %s\n", code, text)
			return nil
		})
		remote := con.RemoteAddr()
		fmt.Printf("Get a new connection %v\n", remote.String())
		ConMap[remote.String()] = con
		/*
			for {
				msgType, data, err := con.ReadMessage()
				if err != nil {
					fmt.Printf("Next Reader error %v\n", err)
					return
				}
				fmt.Printf("get message type %v data:%v\n", msgType, string(data))
			}
		*/
	})

	server := http.Server{
		Addr:    ":9903",
		Handler: mux,
	}

	server.ListenAndServeTLS("./server.crt", "./server.key")
}
