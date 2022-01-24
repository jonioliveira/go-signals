package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	_, cancel := context.WithCancel(context.Background())

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

		<-exit
		cancel()
	}()

	var wg sync.WaitGroup

	httpServer := &http.Server{
		Addr: ":8000",
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		httpServer.ListenAndServe()
	}()

	wg.Wait()
	fmt.Println("Main done")
}
