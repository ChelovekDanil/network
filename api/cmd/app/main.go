package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ChelovekDanil/network/internal/app"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTSTP, syscall.SIGQUIT)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sysCall := <-signalChan
		fmt.Printf("вызван системный вызов: %+v\n", sysCall)
		cancel()
	}()

	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
