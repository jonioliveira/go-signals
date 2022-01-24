package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

const port = "8080"

func main() {
	// graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())

	// handle stop/kill signal
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	httpsvr := &http.Server{
		Addr: ":" + port,
	}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return httpsvr.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		return httpsvr.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
}
