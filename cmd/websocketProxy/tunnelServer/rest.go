package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type ExecRest struct {
	ctx     context.Context
	session *TunnelServer
}

const (
	RequestTypeExec = "exec"
	RequestTypeLogs = "logs"
)

const (
	MessageTypeConnect = "NEW_CONNECTION"
	MessageTypeData    = "DATA"
)

type message struct {
	ConnectID   uint64
	RequestType string
	MessageType string
	Data        []byte
}

func (e ExecRest) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	host := request.Host
	con, err := Upgrader.Upgrade(writer, request, nil)
	if err != nil {
		return
	}
	defer con.Close()

	fmt.Printf("Get a new exec connection to host %v \n", host)
	remotecon, ok := e.session.GetTunnelCon(host)
	if !ok {
		fmt.Printf("no valid tunnel %v...\n", host)
		return
	}
	for {
		t, reader, err := con.NextReader()
		if err != nil {
			fmt.Printf("get next reader error %v\n", err)
			return
		}

		writer, err := remotecon.NextWriter(t)
		if err != nil {
			fmt.Printf("get nextwriter error %v\n", err)
			return
		}

		if _, err := io.Copy(writer, reader); err != nil {
			fmt.Printf("io copy error %v\n", err)
			return
		}
	}
}

func NewExecRest(ctx context.Context, manager *TunnelServer) *ExecRest {
	return &ExecRest{
		ctx:     ctx,
		session: manager,
	}
}
