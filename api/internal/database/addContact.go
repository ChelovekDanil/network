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

type AddContactStore struct{}

func NewAddContactStore() *AddContactStore {
	return &AddContactStore{}
}

func (s *AddContactStore) AddContact(ctx context.Context, id, addLogin string) error {
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

func (s *AddContactStore) isBusy(ctx context.Context, id, addLogin string) error {
	ctx, cancel := context.WithTimeout(ctx, duringSqlQuery)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT id FROM contact WHERE login = $1 AND user_id = $2", addLogin, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil
	}
	return ErrContactBusy
}
