package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var cloudAddr string

func init() {
	flag.StringVar(&cloudAddr, "cloudaddr", cloudAddr, "cloudaddr such as 10.20.10.9:8091")
}

func main() {
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go startLocalServer(ctx)
	for {
		time.Sleep(time.Hour)
	}
}

func startLocalClient(ctx context.Context) net.Conn {
	var (
		temporary time.Duration
		err       error
		con       net.Conn
	)

	for {
		con, err = net.Dial("tcp4", cloudAddr)
		if err == nil {
			break
		}
		if temporary == 0 {
			temporary = time.Millisecond * 50
		} else {
			temporary *= 2
		}
		if max := time.Second * 2; temporary > max {
			temporary = max
		}
		fmt.Printf("dial error %v\n", err)
		time.Sleep(temporary)
		continue
	}
	return con
}

func startLocalServer(ctx context.Context) {
	//http.ListenAndServe()
	var temporary time.Duration
	localCon := startLocalClient(ctx)
	l, err := net.Listen("tcp4", ":8090")
	if err != nil {
		log.Fatalf("listen error %v", err)
	}
	fmt.Println("listen at 8090 ...")
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
		go Server(con, localCon)
	}
}

func Server(con, localCon net.Conn) {
	fmt.Printf("new connection remote addr %s\n", con.RemoteAddr().String())
	for {
		con.SetDeadline(time.Now().Add(time.Second * 2))
		io.Copy(localCon, con)
		/*
			bread := bufio.NewReader(con)
			//io.Copy()
			s, err := bread.ReadString('\n')
			if err != nil {
				if ne, ok := err.(net.Error); ok {
					if ne.Timeout() {
						log.Printf("read string timeout\n")
						continue
					}
				}
				log.Printf("read string error %s", err)
			}

			fmt.Printf("read string:%s\n", s)
		*/
	}
}
