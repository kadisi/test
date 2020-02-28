package main

import (
	"context"
	"fmt"
	"time"
)

func test(rootCtx context.Context) {
	_, cancel := context.WithCancel(rootCtx)
	defer cancel()
	time.Sleep(time.Second * 2)
	fmt.Println("test end ...")
}
func main() {
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	go test(ctx)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("main get Done ..")
			return
		default:
			fmt.Println("main can not get Done")
			time.Sleep(time.Second)
		}
	}
}
