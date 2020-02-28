package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	url2 "net/url"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	url := url2.URL{
		Scheme: "wss",
		Host:   "localhost:9904",
		Path:   "/exec",
	}

	con, _, err := dialer.DialContext(ctx, url.String(), nil)
	defer con.Close()

	go func(ctx context.Context) {
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

	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}
