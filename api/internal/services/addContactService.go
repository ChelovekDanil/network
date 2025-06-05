package services

import (
	"context"
)

type AddContactService struct {
	usrStr userStore
	actStr addContactStore
}

type addContactStore interface {
	AddContact(ctx context.Context, id, addLogin string) error
}

func NewAddContactService(uStr userStore, aStr addContactStore) *AddContactService {
	return &AddContactService{
		usrStr: uStr,
		actStr: aStr,
	}
}

func (s *AddContactService) AddContact(ctx context.Context, login, addLogin string) error {
	u, err := s.usrStr.Get(ctx, login)
	if err != nil {
		return err
	}

	_, err = s.usrStr.Get(ctx, addLogin)
	if err != nil {
		return err
	}

	return s.actStr.AddContact(ctx, u.ID, addLogin)
}
