package database

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
)

var (
	ErrContactBusy = errors.New("контакт уже существует")
)

type ContactStore struct{}

func NewContactStore() *ContactStore {
	return &ContactStore{}
}

func (s *ContactStore) AddContact(ctx context.Context, id, addLogin string) error {
	err := s.isBusy(ctx, id, addLogin)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, duringSqlQuery)
	defer cancel()

	uID := uuid.New()

	if _, err := db.ExecContext(ctx,
		"INSERT INTO contact(id, login, user_id) VALUES($1, $2, $3);",
		uID,
		addLogin,
		id,
	); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *ContactStore) DeleteContact(ctx context.Context, id, addLogin string) error {
	ctx, cancel := context.WithTimeout(ctx, duringSqlQuery)
	defer cancel()

	if _, err := db.ExecContext(ctx,
		"DELETE FROM contact WHERE login = $1 AND user_id = $2",
		addLogin,
		id,
	); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *ContactStore) GetAll(ctx context.Context, id string) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, duringSqlQuery)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT login FROM contact WHERE user_id = $1;", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]string, 0)

	for rows.Next() {
		var user string
		if err := rows.Scan(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (s *ContactStore) Message(ctx context.Context, firstID, lastID, message string) error {
	ctx, cancel := context.WithTimeout(ctx, duringSqlQuery)
	defer cancel()

	uID := uuid.New()

	if _, err := db.ExecContext(ctx,
		"INSERT INTO message(id, first_login, last_login, message) VALUES($1, $2, $3, $4);",
		uID,
		firstID,
		lastID,
		message,
	); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *ContactStore) GetMessage(ctx context.Context, firstId, lastID string) ([][]string, error) {
	ctx, cancel := context.WithTimeout(ctx, duringSqlQuery)
	defer cancel()

	first := []string{}
	last := []string{}

	rows, err := db.QueryContext(ctx, "SELECT message FROM message WHERE first_login = $1 AND last_login = $2", firstId, lastID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var message string
		if err := rows.Scan(&message); err != nil {
			return nil, err
		}
		first = append(first, message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	rows, err = db.QueryContext(ctx, "SELECT message FROM message WHERE first_login = $1 AND last_login = $2;", lastID, firstId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var message string
		if err := rows.Scan(&message); err != nil {
			return nil, err
		}
		first = append(first, message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return [][]string{first, last}, nil
}

func (s *ContactStore) isBusy(ctx context.Context, id, addLogin string) error {
	ctx, cancel := context.WithTimeout(ctx, duringSqlQuery)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT message FROM message WHERE first_login = $1 AND last_login = $2;", addLogin, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil
	}
	return ErrContactBusy
}
