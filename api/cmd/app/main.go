package main

import (
	"log"

	"github.com/ChelovekDanil/network/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
