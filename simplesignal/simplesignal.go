package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Signal(0xf)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		fmt.Printf("\n%s", sig)
		done <- true
	}()
	fmt.Println("awaiting signal")
	<-done
}
