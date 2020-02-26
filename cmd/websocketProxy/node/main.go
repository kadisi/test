package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	url2 "net/url"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	url := url2.URL{
		Scheme: "wss",
		Host:   "localhost:9903",
		Path:   "/test",
	}
	pool := x509.NewCertPool()
	cadate, err := ioutil.ReadFile("./ca.crt")
	if err != nil {
		fmt.Printf("read ca file error %v\n", err)
		return
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
		return
	}
	defer con.Close()

	for {
		err := con.WriteMessage(websocket.TextMessage, []byte("hello world\n"))
		if err != nil {
			fmt.Printf("write message error %v\n", err)
		}
		time.Sleep(time.Second)
	}

}
