package services

import (
	"context"
	"fmt"
)

type ContactService struct {
	usrStr userStore
	actStr contactStore
}

type contactStore interface {
	AddContact(ctx context.Context, id, addLogin string) error
	DeleteContact(ctx context.Context, id, addLogin string) error
	GetAll(ctx context.Context, login string) ([]string, error)
	Message(ctx context.Context, firstLogin, lastLogin, message string) error
	GetMessage(ctx context.Context, firstLogin, lastLogin string) ([][]string, error)
}

func NewContactService(uStr userStore, aStr contactStore) *ContactService {
	return &ContactService{
		usrStr: uStr,
		actStr: aStr,
	}
}

func (s *ContactService) AddContact(ctx context.Context, login, addLogin string) error {
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

func (s *ContactService) DeleteContact(ctx context.Context, login, addLogin string) error {
	u, err := s.usrStr.Get(ctx, login)
	if err != nil {
		return err
	}
	return s.actStr.DeleteContact(ctx, u.ID, addLogin)
}

func (s *ContactService) GetAll(ctx context.Context, login string) ([]string, error) {
	u, err := s.usrStr.Get(ctx, login)
	if err != nil {
		return nil, err
	}

	return s.actStr.GetAll(ctx, u.ID)
}

func (s *ContactService) Message(ctx context.Context, firstLogin, lastLogin, message string) error {
	u1, err := s.usrStr.Get(ctx, firstLogin)
	if err != nil {
		return err
	}
	u2, err := s.usrStr.Get(ctx, lastLogin)
	if err != nil {
		return err
	}
	if message == "" {
		return fmt.Errorf("empty message")
	}

	return s.actStr.Message(ctx, u1.ID, u2.ID, message)
}

func (s *ContactService) GetMessage(ctx context.Context, firstLogin, lastLogin string) ([][]string, error) {
	u1, err := s.usrStr.Get(ctx, firstLogin)
	if err != nil {
		return nil, err
	}

	u2, err := s.usrStr.Get(ctx, lastLogin)
	if err != nil {
		return nil, err
	}

	return s.actStr.GetMessage(ctx, u1.ID, u2.ID)
}
