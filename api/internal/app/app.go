package app

import (
	"github.com/ChelovekDanil/network/internal/database"
	"github.com/ChelovekDanil/network/internal/transport/rest"
)

func Run() error {
	if err := database.InitDB(); err != nil {
		return err
	}
	rest.Start()
	return nil
}
