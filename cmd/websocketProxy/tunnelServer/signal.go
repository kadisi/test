package main

import (
	"os"
	"os/signal"
	"syscall"
)

func signalWatch(shutdown func()) {
	s := make(chan os.Signal, 10)
	signal.Notify(s, syscall.SIGHUP, syscall.SIGINT)

	go func() {
		for {
			select {
			case <-s:
				shutdown()

				return
			}
		}
	}()

}
