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
	ErrUserExist     = errors.New("user exist")
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
	u, err := s.userStr.Get(ctx, user.Login)
	if err != nil {
		return nil, err
	}

	if u.PassHash != cryptocs.Hash(user.PassHash) {
		return nil, ErrWrongPassword
	}

	accessToken, err := genAccessToken(u.Login)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.tokenStr.CreateToken(ctx, u.ID)
	if err != nil {
		return nil, err
	}

	return []string{accessToken, refreshToken}, nil
}

func (s *AuthService) ReLogin(ctx context.Context, login string, refreshToken string) ([]string, error) {
	u, err := s.userStr.Get(ctx, login)
	if err != nil {
		return nil, err
	}

	reToken, err := s.tokenStr.RefreshToken(ctx, u.ID, refreshToken)
	if err != nil {
		return nil, err
	}

	accessToken, err := genAccessToken(login)
	if err != nil {
		return nil, err
	}

	return []string{accessToken, reToken}, nil
}

func (s *AuthService) Register(ctx context.Context, user models.User) ([]string, error) {
	u, err := s.userStr.Get(ctx, user.ID)
	if err == nil && u != nil {
		return nil, ErrUserExist
	}

	id, err := s.userStr.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	accessToken, err := genAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.tokenStr.CreateToken(ctx, id)
	if err != nil {
		return nil, err
	}

	return []string{accessToken, refreshToken}, err
}

// генерирует access jwt токен
func genAccessToken(data string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"data": data,
			"exp":  time.Now().Add(time.Hour * 100).Unix(),
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
