package services

import (
	"context"

	"github.com/ChelovekDanil/network/internal/models"
)

type UserService struct {
	store userStore
}

type userStore interface {
	Get(ctx context.Context, login string) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	Create(ctx context.Context, user models.User) (string, error)
	Update(ctx context.Context, login string, user models.User) error
	Delete(ctx context.Context, login string) error
}

func NewUserService(s userStore) *UserService {
	return &UserService{
		store: s,
	}
}

// Get возвращает пользователя по login
func (s *UserService) Get(ctx context.Context, login string) (*models.User, error) {
	return s.store.Get(ctx, login)
}

// GetAll возвращает всех пользователей
func (s *UserService) GetAll(ctx context.Context) ([]models.User, error) {
	return s.store.GetAll(ctx)
}

// Create создает и сохраняет нового пользователя
func (s *UserService) Create(ctx context.Context, user models.User) (string, error) {
	return s.store.Create(ctx, user)
}

// Update обновляет пользователя по login
func (s *UserService) Update(ctx context.Context, login string, user models.User) error {
	return s.store.Update(ctx, login, user)
}

// Delete удаляет пользователя по login
func (s *UserService) Delete(ctx context.Context, login string) error {
	return s.store.Delete(ctx, login)
}
