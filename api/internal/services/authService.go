package services

import (
	"context"
	"errors"
	"time"

	"github.com/ChelovekDanil/network/internal/lib/cryptocs"
	"github.com/ChelovekDanil/network/internal/models"
	"github.com/golang-jwt/jwt"
)

var (
	secretKey        = []byte("scretni-lol")
	ErrWrongPassword = errors.New("wrong password")
)

type authStore interface {
	CreateToken(ctx context.Context, userID string) (string, error)
	RefreshToken(ctx context.Context, userID string, refreshToken string) (string, error)
}

type AuthService struct {
	userStr  userStore
	tokenStr authStore
}

func NewAuthService(uStr userStore, tStr authStore) *AuthService {
	return &AuthService{
		userStr:  uStr,
		tokenStr: tStr,
	}
}

func (s *AuthService) Login(ctx context.Context, user models.User) ([]string, error) {
	u, err := s.userStr.Get(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	if u.PassHash != cryptocs.Hash(user.PassHash) {
		return nil, ErrWrongPassword
	}

	accessToken, err := genAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.tokenStr.CreateToken(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return []string{accessToken, refreshToken}, nil
}

func (s *AuthService) ReLogin(ctx context.Context, user models.User, refreshToken string) ([]string, error) {
	_, err := s.userStr.Get(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	reToken, err := s.tokenStr.RefreshToken(ctx, user.ID, refreshToken)
	if err != nil {
		return nil, err
	}

	accessToken, err := genAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	return []string{accessToken, reToken}, nil
}

func genAccessToken(data string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"data": data,
			"exp":  time.Now().Add(time.Minute * 10).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetSecKey() []byte {
	return secretKey
}
