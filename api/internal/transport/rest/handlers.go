package rest

import (
	"net/http"

	"github.com/ChelovekDanil/network/internal/database"
	"github.com/ChelovekDanil/network/internal/services"
)

func Start() error {
	userStore := database.NewUserStore()
	userService := services.NewUserService(userStore)
	userHandler := NewUserHandler(userService)

	mux := http.NewServeMux()
	mux.Handle("/user/", userHandler)

	return http.ListenAndServe(":8080", mux)
}
