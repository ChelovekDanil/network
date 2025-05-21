package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/ChelovekDanil/network/internal/lib/cryptocs"
	"github.com/ChelovekDanil/network/internal/models"
	"github.com/google/uuid"
)

var (
	ErrUserNotFound  = errors.New("user not found")     // пользователь не найден
	ErrUserLoginBusy = errors.New("user login is busy") // логин пользователя занят
	duringSqlQuery   = time.Second * 2                  // максимальное время запроса
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

	rows, err := db.QueryContext(ctx, "SELECT id, login, passhash FROM users WHERE id = $1;", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, ErrUserNotFound
	}

	var user models.User
	if err := rows.Scan(&user.ID, &user.Login, &user.PassHash); err != nil {
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

	rows, err := db.QueryContext(ctx, "SELECT id, login FROM users;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Login); err != nil {
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

	isBusy, err := isBusyLogin(ctx, user.Login)
	if err != nil {
		log.Println(err)
		return err
	}
	if isBusy {
		log.Println(ErrUserLoginBusy.Error())
		return ErrUserLoginBusy
	}

	hashpass := cryptocs.Hash(user.PassHash)
	log.Println(hashpass)
	uID := uuid.New()

	if _, err := db.ExecContext(ctx,
		"INSERT INTO users(id, login, passhash) VALUES($1, $2, $3);",
		uID,
		user.Login,
		hashpass,
	); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// Update обновляет пользователя из бд
func (s *UserStore) Update(ctx context.Context, id string, user models.User) error {
	return nil

	// ctx, cancel := context.WithTimeout(ctx, duringSqlQuery)
	// defer cancel()

	// if _, err := db.ExecContext(ctx,
	// 	"UPDATE users SET id = $1, login = $2, passhash = $3 WHERE id = $4;",
	// 	user.ID,
	// 	user.Login,
	// 	user.PassHash,
	// 	id,
	// ); err != nil {
	// 	return err
	// }
	// return nil
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

func isBusyLogin(ctx context.Context, login string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, duringSqlQuery)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT id, login FROM users;")
	if err != nil {
		return true, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Login); err != nil {
			return true, err
		}
		if user.Login == login {
			return true, nil
		}
	}

	if err := rows.Err(); err != nil {
		return true, err
	}

	return false, nil
}
