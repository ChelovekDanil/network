package database

import (
	"context"
	"errors"

	"github.com/ChelovekDanil/network/internal/models"
)

var (
	UserNotFoundError = errors.New("user not found")
)

type UserStore struct {
}

func NewUserStore() *UserStore {
	return &UserStore{}
}

func (s *UserStore) Get(id string) (*models.User, error) {
	rows, err := db.QueryContext(context.Background(), "SELECT id, firstname, lastname FROM users WHERE id = $1;", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, UserNotFoundError
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

func (s *UserStore) GetAll() ([]models.User, error) {
	rows, err := db.QueryContext(context.Background(), "SELECT id, firstname, lastname FROM users;")
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

func (s *UserStore) Create(user models.User) error {
	if _, err := db.ExecContext(context.Background(),
		"INSERT INTO users(id, firstname, lastname) VALUES($1, $2, $3);",
		user.ID,
		user.FirstName,
		user.LastName,
	); err != nil {
		return err
	}

	return nil
}

func (s *UserStore) Update(id string, user models.User) error {
	if _, err := db.ExecContext(context.Background(),
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

func (s *UserStore) Delete(id string) error {
	if _, err := db.ExecContext(context.Background(),
		"DELETE FROM users WHERE id = $1;",
		id,
	); err != nil {
		return err
	}
	return nil
}
