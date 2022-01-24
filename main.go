package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	defer os.Exit(0)

	_, cancel := context.WithCancel(context.Background())
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)

	cancel()
}
