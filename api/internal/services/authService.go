package services

import "github.com/ChelovekDanil/network/internal/models"

type AuthService struct {
	userStr  userStore
	tokenStr tokenStore
}

type tokenStore interface {
	GetToken(userID string) (string, error)
	CreteToken(userID string, token string) error
}

func (s *AuthService) Login(userID string, token string) error {

	return nil
}

func (s *AuthService) Register(user models.User, token string) error {
	return nil
}
