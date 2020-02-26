package main

import (
	"fmt"
	url2 "net/url"
	"path"
)

func main() {
	server := "127.0.0.1:1111"
	projectid := "12344"
	nodeid := "56789"
	url := url2.URL{
		Scheme: "wss",
		Host:   server,
		Path:   path.Join(projectid, nodeid, "events"),
	}
	path.Join()
	fmt.Println(url.String())
}
