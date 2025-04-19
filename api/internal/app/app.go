package app

import (
	"context"
	"time"

	"github.com/ChelovekDanil/network/internal/database"
	"github.com/ChelovekDanil/network/internal/transport/rest"
)

const (
	timeRequestDB = 5 * time.Second
)

func Run(ctx context.Context) error {
	ctxTime, cancel := context.WithTimeout(context.Background(), timeRequestDB)
	defer cancel()

	errChan := make(chan error)
	go func(err chan<- error) {
		err <- database.InitDB(ctxTime)
	}(errChan)

	err := <-errChan
	if err != nil {
		return err
	}

	go func() {
		rest.Start(ctx)
	}()

	<-ctx.Done()
	return nil
}
