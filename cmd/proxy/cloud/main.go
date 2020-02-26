package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go startLocalServer(ctx)
	for {
		time.Sleep(time.Hour)
	}
}

func startLocalServer(ctx context.Context) {
	//http.ListenAndServe()
	var temporary time.Duration
	l, err := net.Listen("tcp4", ":8091")
	if err != nil {
		log.Fatalf("listen error %v", err)
	}
	fmt.Println("listen at 8091 ...")
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		con, err := l.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if temporary == 0 {
					temporary = 5 * time.Millisecond
				} else {
					temporary *= 2
				}
				if temporary > time.Second {
					temporary = time.Second
				}
				log.Printf("http: Accept error: %v; retrying in %v", err, temporary)
				time.Sleep(temporary)
				continue
			}
			return
		}
		fmt.Printf("accept on connection %v\n", con.RemoteAddr().String())
		go Server(con)
	}
}

func Server(con net.Conn) {
	fmt.Printf("cloud new connection remote addr %s\n", con.RemoteAddr().String())
	for {
		con.SetDeadline(time.Now().Add(time.Second * 2))
		bread := bufio.NewReader(con)
		//io.Copy()
		s, err := bread.ReadString('\n')
		if err != nil {
			if ne, ok := err.(net.Error); ok {
				if ne.Timeout() {
					continue
				}
			}
			log.Printf("read string error %s", err)
		}

		fmt.Printf("cloud read string:%s\n", s)
	}
}
