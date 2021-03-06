package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		w.WriteHeader(status)
		w.Write([]byte(reason.Error()))
	},
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sessionManager := NewTunnelServer(ctx)
	execRest := NewExecRest(ctx, sessionManager)

	mux := http.NewServeMux()
	mux.Handle("/connect", sessionManager)
	mux.Handle("/exec", execRest)

	server := &http.Server{
		Addr:    ":10250",
		Handler: mux,
	}

	signalWatch(func() {
		cancel()
		ctx, serverCancel := context.WithTimeout(context.Background(), time.Second*5)
		defer serverCancel()
		server.Shutdown(ctx)
		time.Sleep(time.Second * 5)
		fmt.Println("shutdown succefully ...")
	})

	err := server.ListenAndServeTLS("./server.crt", "./server.key")
	if err != nil {
		log.Fatalf("start server error %v\n", err)
	}
}
