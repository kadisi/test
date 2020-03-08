package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var dest string

func init() {
	flag.StringVar(&dest, "dest", "127.0.0.1:10250", "dest address <ip:port>")
}

func main() {
	flag.Parse()
	certPools := x509.NewCertPool()
	caData, err := ioutil.ReadFile("./ca.crt")
	if err != nil {
		log.Fatal("read ca file error %v\n", err)
	}
	certPools.AppendCertsFromPEM(caData)
	dialer := websocket.Dialer{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
			RootCAs:            certPools,
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	url := url2.URL{
		Scheme: "wss",
		Host:   dest,
		Path:   "/exec",
	}

	header := http.Header{}
	header.Add("HOST", "127.0.0.1")
	con, _, err := dialer.DialContext(ctx, url.String(), header)
	if err != nil {
		log.Fatalf("dail %v error %v", url.String(), err)
	}
	defer con.Close()

	group := sync.WaitGroup{}
	group.Add(1)
	go func(ctx context.Context) {
		defer group.Done()
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Second)
			}
			err := con.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("say hello %v", i)))
			if err != nil {
				fmt.Printf("write message error %v\n", err)
				return
			}
		}
	}(ctx)

	group.Wait()
}
