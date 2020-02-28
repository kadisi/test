package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	url2 "net/url"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

type WSConnect struct {
	Con *websocket.Conn
}

type TunnelSession struct {
	wsCon *WSConnect
}

func NewTunnelSession(c *websocket.Conn) *TunnelSession {
	return &TunnelSession{
		wsCon: &WSConnect{Con: c},
	}
}

func (s *TunnelSession) startPing(ctx context.Context) {
	t := time.NewTicker(time.Second * 5)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			err := s.wsCon.Con.WriteControl(websocket.PingMessage, []byte("ping you"), time.Now().Add(time.Second))
			if err != nil {
				fmt.Printf("write ping message error %v\n", err)
			}
		}
	}

}
func (s *TunnelSession) Serve() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go s.startPing(ctx)

	for {
		t, data, err := s.wsCon.Con.ReadMessage()
		if err != nil {
			fmt.Printf("Read Message error %v\n")
			return err
		}
		fmt.Printf("receive message type %v data:%v\n", t, data)
	}

	return nil
}

func TLSClientConnect(url url2.URL, ca string) error {
	fmt.Println("start a new connection ...")
	pool := x509.NewCertPool()
	cadate, err := ioutil.ReadFile(ca)
	if err != nil {
		fmt.Printf("read ca file error %v\n", err)
		return err
	}
	pool.AppendCertsFromPEM(cadate)
	dial := websocket.Dialer{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
			RootCAs:            pool,
		},
	}

	con, _, err := dial.Dial(url.String(), nil)
	if err != nil {
		fmt.Printf("dial error %v\n", err)
		return err
	}
	defer con.Close()

	session := NewTunnelSession(con)
	return session.Serve()
}

func main() {
	url := url2.URL{
		Scheme: "wss",
		Host:   "localhost:9904",
		Path:   "/connect",
	}

	d, _ := os.Getwd()
	fmt.Printf("get wd %v\n", d)
	for range time.NewTicker(time.Second).C {
		err := TLSClientConnect(url, "./ca.crt")
		if err != nil {
			fmt.Printf("TLSClientConnect error %v\n", err)
		}
	}
}
