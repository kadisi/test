package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	pool := x509.NewCertPool()
	data, err := ioutil.ReadFile("./ca.crt")
	if err != nil {
		log.Fatalf("read file error %v\n", err)
	}
	pool.AppendCertsFromPEM(data)

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
				RootCAs:            pool,
			},
		},
	}
	req, err := http.NewRequest("GET", "https://localhost:8809/watch", nil)
	if err != nil {
		log.Fatalf("request error %v", err)
	}
	req.Header.Set("Transfer-Encoding", "chunked")
	rep, err := client.Do(req)
	if err != nil {
		log.Fatalf("respon error %v", err)
	}
	scan := bufio.NewScanner(rep.Body)
	for scan.Scan() {
		fmt.Println(scan.Text())
	}
}
