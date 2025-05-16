package rest

import "github.com/ChelovekDanil/network/internal/models"

type AuthHandler struct {
	service authService
}

type authService interface {
	Login(userID string, token string) error
	Register(user models.User, token string) error
}
