package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"time"
)

func main() {
	jar, _ := cookiejar.New(nil)
	client := http.Client{
		Jar: jar,
	}
	req, err := http.NewRequest("GET", "http://localhost:9999/rest/check", nil)
	if err != nil {
		log.Fatalf("create request error %v", err)
	}
	ticker := time.Tick(time.Second)
	for range ticker {
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("client do error %v", err)
		}
		defer resp.Body.Close()
		cookies := resp.Cookies()
		for _, c := range cookies {
			fmt.Printf("get recive cookies: %v\n", c.String())
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("read body error %v", err)
		}
		fmt.Println(string(data))
	}

}
