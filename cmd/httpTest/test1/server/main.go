package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type WatchData struct {
	Second int
	Name   string
}

func (w *WatchData) Byte() ([]byte, error) {
	return json.Marshal(w)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/watch", func(w http.ResponseWriter, request *http.Request) {
		fmt.Printf("get new request from %v\n", request.RemoteAddr)

		flush, ok := w.(http.Flusher)
		if !ok {
			log.Fatalf("change to fluster error\n")
		}
		stop := make(chan struct{})
		go func(stop chan<- struct{}) {
			t := time.NewTicker(time.Second)
			count := 1
			for c := range t.C {
				if count >= 3 {
					break
				}
				fmt.Println("haha prepare to write ...")
				d := &WatchData{
					Second: c.Second(),
					Name:   fmt.Sprintf("zhangejie%v", c.Second()),
				}
				data, err := d.Byte()
				if err != nil {
					log.Fatalf("get byte error %v", err)
				}
				fmt.Fprintf(w, "%v\n", string(data))
				flush.Flush()
				fmt.Println("haha write success...\n")
				count++
			}
			t.Stop()
			stop <- struct{}{}
		}(stop)
		<-stop
		fmt.Println("time sleep end ...")
	})

	server := http.Server{
		Addr:    ":8809",
		Handler: mux,
	}
	server.ListenAndServeTLS("./tunnelServer.crt", "./tunnelServer.key")
}
