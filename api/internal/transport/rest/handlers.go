package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ChelovekDanil/network/internal/database"
	"github.com/ChelovekDanil/network/internal/services"
	"github.com/golang-jwt/jwt"
)

var (
	duringResponse = time.Second * 2 // максимальное время запроса
	ErrNotFoundId  = errors.New("not found id")
)

func Start(ctx context.Context) error {
	userHandler := createUserHandler()
	authHandler := createAuthHandler()

	mux := http.NewServeMux()
	mux.Handle("/user/", userHandler)
	mux.Handle("/auth/", authHandler)

	fmt.Println("server start on port :8080")
	return http.ListenAndServe(":8080", mux)
}

func createUserHandler() *UserHandler {
	userStore := database.NewUserStore()
	userService := services.NewUserService(userStore)
	return NewUserHandler(userService)
}

func createAuthHandler() *AuthHandler {
	userStore := database.NewUserStore()
	authStore := database.NewAuthStore()
	authService := services.NewAuthService(userStore, authStore)
	return NewAuthHandler(authService)
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(toekn *jwt.Token) (any, error) {
		return services.GetSecKey(), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}

func protectedHandler(w http.ResponseWriter, r *http.Request) bool {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return false
	}
	tokenString = tokenString[len("Bearer "):]

	err := verifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return false
	}

	return true
}
