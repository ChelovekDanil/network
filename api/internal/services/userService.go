package services

import "github.com/ChelovekDanil/network/internal/models"

type UserService struct {
	store userStore
}

type userStore interface {
	Get(id string) (*models.User, error)
	GetAll() ([]models.User, error)
	Create(user models.User) error
	Update(id string, user models.User) error
	Delete(id string) error
}

func NewUserService(s userStore) *UserService {
	return &UserService{
		store: s,
	}
}

func (s *UserService) Get(id string) (*models.User, error) {
	return s.store.Get(id)
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.store.GetAll()
}

func (s *UserService) Create(user models.User) error {
	return s.store.Create(user)
}

func (s *UserService) Update(id string, user models.User) error {
	return s.store.Update(id, user)
}

func (s *UserService) Delete(id string) error {
	return s.store.Delete(id)
}
