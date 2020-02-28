package main

import (
	"context"
	"fmt"
	"net/http"
)

type ExecRest struct {
	ctx context.Context
}

func (e ExecRest) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	con, err := Upgrader.Upgrade(writer, request, nil)
	if err != nil {
		return
	}
	defer con.Close()
	remote := con.RemoteAddr().String()
	local := con.LocalAddr().String()
	fmt.Printf("Get a new exec connection remote %v local %v\n",
		remote, local)
	for {
		_, data, err := con.ReadMessage()
		if err != nil {
			fmt.Printf("readmessage error %v\n", err)
			return
		}
		fmt.Printf("receive message %v from %v to %v\n", string(data), remote, local)
	}
}

func NewExecRest(ctx context.Context) *ExecRest {
	return &ExecRest{
		ctx: ctx,
	}
}
