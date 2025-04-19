package database

import (
	"context"
	"errors"
	"time"

	"github.com/ChelovekDanil/network/internal/models"
)

var (
	ErrUserNotFound = errors.New("user not found") // пользователь не найден
	duringSqlQuery  = time.Second * 2              // максимальное время запроса
)

// UserStore реализует операции взаимодействие с бд для модели пользователя
type UserStore struct {
}

func NewUserStore() *UserStore {
	return &UserStore{}
}

// Get достает модель пользователя из базы данных
func (s *UserStore) Get(ctx context.Context, id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, duringSqlQuery)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT id, firstname, lastname FROM users WHERE id = $1;", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, ErrUserNotFound
	}

	var user models.User
	if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetAll возвращает всех пользователей из бд
func (s *UserStore) GetAll(ctx context.Context) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, duringSqlQuery)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT id, firstname, lastname FROM users;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

// Create создает нового пользователя и сохраняет в бд
func (s *UserStore) Create(ctx context.Context, user models.User) error {
	ctx, cancel := context.WithTimeout(ctx, duringSqlQuery)
	defer cancel()

	if _, err := db.ExecContext(ctx,
		"INSERT INTO users(id, firstname, lastname) VALUES($1, $2, $3);",
		user.ID,
		user.FirstName,
		user.LastName,
	); err != nil {
		return err
	}

	return nil
}

// Update обновляет пользователя из бд
func (s *UserStore) Update(ctx context.Context, id string, user models.User) error {
	ctx, cancel := context.WithTimeout(ctx, duringSqlQuery)
	defer cancel()

	if _, err := db.ExecContext(ctx,
		"UPDATE users SET id = $1, firstname = $2, lastname = $3 WHERE id = $4;",
		user.ID,
		user.FirstName,
		user.LastName,
		id,
	); err != nil {
		return err
	}
	return nil
}

// Delete удаляет пользователя из бд
func (s *UserStore) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, duringSqlQuery)
	defer cancel()

	if _, err := db.ExecContext(ctx,
		"DELETE FROM users WHERE id = $1;",
		id,
	); err != nil {
		return err
	}
	return nil
}
