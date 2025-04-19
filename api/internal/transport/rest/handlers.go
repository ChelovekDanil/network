package rest

import (
	"context"
	"net/http"

	"github.com/ChelovekDanil/network/internal/database"
	"github.com/ChelovekDanil/network/internal/services"
)

func Start(ctx context.Context) error {
	userHandler := createUserHandler()

	mux := http.NewServeMux()
	mux.Handle("/user/", userHandler)
	return http.ListenAndServe(":8080", mux)
}

func createUserHandler() *UserHandler {
	userStore := database.NewUserStore()
	userService := services.NewUserService(userStore)
	return NewUserHandler(userService)
}
