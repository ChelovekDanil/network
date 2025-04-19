package services

import (
	"context"

	"github.com/ChelovekDanil/network/internal/models"
)

type UserService struct {
	store userStore
}

type userStore interface {
	Get(ctx context.Context, id string) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	Create(ctx context.Context, user models.User) error
	Update(ctx context.Context, id string, user models.User) error
	Delete(ctx context.Context, id string) error
}

func NewUserService(s userStore) *UserService {
	return &UserService{
		store: s,
	}
}

// Get возвращает пользователя по id
func (s *UserService) Get(ctx context.Context, id string) (*models.User, error) {
	return s.store.Get(ctx, id)
}

// GetAll возвращает всех пользователей
func (s *UserService) GetAll(ctx context.Context) ([]models.User, error) {
	return s.store.GetAll(ctx)
}

// Create создает и сохраняет нового пользователя
func (s *UserService) Create(ctx context.Context, user models.User) error {
	return s.store.Create(ctx, user)
}

// Update обновляет пользователя по id
func (s *UserService) Update(ctx context.Context, id string, user models.User) error {
	return s.store.Update(ctx, id, user)
}

// Delete удаляет пользователя по id
func (s *UserService) Delete(ctx context.Context, id string) error {
	return s.store.Delete(ctx, id)
}
