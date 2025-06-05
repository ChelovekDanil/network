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
	ErrNotFoundId  = errors.New("not found login")
)

func Start(ctx context.Context) error {
	userHandler := createUserHandler()
	authHandler := createAuthHandler()
	addContactHandler := createContactHandler()

	mux := http.NewServeMux()
	// /user/{login}/ GET  	  - получение пользователя
	// /user/ 	      GET  	  - получение всех пользователей
	// /user/ 	      POST 	  - создание пользователя
	// /user/{login}/ PUT  	  - редактирование пользователя
	// /user/{login}/ DELETE  - удаление пользователя
	mux.Handle("/user/", corsMiddleware(userHandler))
	// /auth/login/    POST - получение access и refresh токена | login, passhash
	// /auth/refresh/  POST - получение access и нового refresh токена | login, refreshToken
	// /auth/register/ POST - создание пользователя и получение access и refresh токена | login, passhash
	mux.Handle("/auth/", corsMiddleware(authHandler))
	// /contact/add/        POST - добавление нового контакта к пользователю | firstLogin, lastLogin
	// /contact/delete/     POST - удаление контакта | firstLogin, lastLogin
	// /contact/{login} GET  - получить контакты пользователя
	// /contact/message/    POST - отправка сообщения | firstLogin, lastLogin, message
	mux.Handle("/contact/", corsMiddleware(addContactHandler))

	fmt.Println("server start on port :8080")
	return http.ListenAndServe(":8080", mux)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
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

func createContactHandler() *ContactHandler {
	userStore := database.NewUserStore()
	addContactStore := database.NewContactStore()
	addContactService := services.NewContactService(userStore, addContactStore)
	return NewContactHanlder(addContactService)
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
