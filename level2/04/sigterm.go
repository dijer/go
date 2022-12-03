package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancelSignal := context.WithCancel(context.Background())
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	defer cancelSignal()

	signal := <-signals
	fmt.Println("Received SIGTERM", signal)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	<-ctx.Done()
	fmt.Println("Closed")
}
