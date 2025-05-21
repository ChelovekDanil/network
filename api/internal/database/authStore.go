package database

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
)

var (
	ErrTokenNotFound = errors.New("token not found")
)

type AuthStore struct {
}

func NewAuthStore() *AuthStore {
	return &AuthStore{}
}

func (s *AuthStore) CreateToken(ctx context.Context, userID string) (string, error) {
	refreshToken, err := genAndSaveRefreshToken(ctx, userID)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func (s *AuthStore) RefreshToken(ctx context.Context, userID string, refreshToken string) (string, error) {
	// get old refresh token
	oldRefreshToken, err := getRefreshToken(ctx, userID)
	if err != nil {
		return "", err
	}

	if oldRefreshToken != refreshToken {
		return "", ErrTokenNotFound
	}

	// delete old refresh token
	if err := deleteOldRefreshToken(ctx, userID); err != nil {
		return "", err
	}

	// gen and save new refresh token
	newRefreshToken, err := genAndSaveRefreshToken(ctx, userID)
	if err != nil {
		return "", err
	}

	return newRefreshToken, nil
}

func getRefreshToken(ctx context.Context, userID string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, duringSqlQuery)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT token FROM token WHERE user_id = $1", userID)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if !rows.Next() {
		return "", ErrTokenNotFound
	}

	var reToken string
	if err := rows.Scan(&reToken); err != nil {
		return "", err
	}

	if err := rows.Err(); err != nil {
		return "", err
	}

	return reToken, nil
}

func genAndSaveRefreshToken(ctx context.Context, userID string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, duringSqlQuery)
	defer cancel()

	id := uuid.New().String()
	newRefreshToken := uuid.New().String()

	if _, err := db.ExecContext(ctx,
		"INSERT INTO token(id, user_id, token) VALUES($1, $2, $3)",
		id,
		userID,
		newRefreshToken,
	); err != nil {
		log.Println(err)
		return "", err
	}

	return newRefreshToken, nil
}

func deleteOldRefreshToken(ctx context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, duringSqlQuery)
	defer cancel()

	if _, err := db.ExecContext(ctx, "DELETE FROM token WHERE user_id = $1", userID); err != nil {
		return err
	}

	return nil
}
