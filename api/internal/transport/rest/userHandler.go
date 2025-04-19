package rest

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/ChelovekDanil/network/internal/models"
)

var (
	userRe            = regexp.MustCompile(`^/user/$`)
	userReSingleParam = regexp.MustCompile(`^/user/.*/$`)
	duringResponse    = time.Second * 2 // максимальное время запроса
)

var (
	ErrNotFoundId = errors.New("not found id")
)

type UserHandler struct {
	service userService
}

type userService interface {
	Get(ctx context.Context, id string) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	Create(ctx context.Context, user models.User) error
	Update(ctx context.Context, id string, user models.User) error
	Delete(ctx context.Context, id string) error
}

func NewUserHandler(s userService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

// ServeHTTP сопоставляет запрос с обработчиком
func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && userReSingleParam.MatchString(r.URL.Path):
		h.GetUser(w, r)
	case r.Method == http.MethodGet && userRe.MatchString(r.URL.Path):
		h.GetAllUsers(w, r)
	case r.Method == http.MethodPost && userRe.MatchString(r.URL.Path):
		h.CreateUser(w, r)
	case r.Method == http.MethodPut && userReSingleParam.MatchString(r.URL.Path):
		h.UpdateUser(w, r)
	case r.Method == http.MethodDelete && userReSingleParam.MatchString(r.URL.Path):
		h.DeleteUser(w, r)
	}
}

// GetUser отправлет модель пользователя
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), duringResponse)
	defer cancel()

	param := getParamFromPath(r.URL.Path)

	user, err := h.service.Get(ctx, param)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	writeResponseOKWithData(w, jsonUser)
}

// GetAllUsers отправляет все существующие модели пользователей
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), duringResponse)
	defer cancel()

	users, err := h.service.GetAll(ctx)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	jsonUsers, err := json.Marshal(users)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	writeResponseOKWithData(w, jsonUsers)
}

// CreateUser создает и сохранет пользователя в бд
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), duringResponse)
	defer cancel()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	if err := h.service.Create(ctx, user); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// UpdateUser обновляет пользователя в бд
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), duringResponse)
	defer cancel()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	id := getParamFromPath(r.URL.Path)
	if err := h.service.Update(ctx, id, user); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteUser удаляет пользователя из бд
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), duringResponse)
	defer cancel()

	id := getParamFromPath(r.URL.Path)
	if err := h.service.Delete(ctx, id); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}

func getParamFromPath(path string) string {
	params := strings.Split(path, "/")
	return params[len(params)-2]
}

func writeResponseOKWithData(w http.ResponseWriter, data []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
