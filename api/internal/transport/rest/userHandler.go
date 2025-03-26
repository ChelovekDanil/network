package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/ChelovekDanil/network/internal/models"
)

var (
	userRe            = regexp.MustCompile(`^/user/$`)
	userReSingleParam = regexp.MustCompile(`^/user/.*/$`)
)

var (
	NotFoundIdError = errors.New("not found id")
)

type UserHandler struct {
	service userService
}

type userService interface {
	Get(id string) (*models.User, error)
	GetAll() ([]models.User, error)
	Create(user models.User) error
	Update(id string, user models.User) error
	Delete(id string) error
}

func NewUserHandler(s userService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

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

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	param := getParamFromPath(r.URL.Path)

	user, err := h.service.Get(param)
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

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAll()
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

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	if err := h.service.Create(user); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	id := getParamFromPath(r.URL.Path)
	if err := h.service.Update(id, user); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := getParamFromPath(r.URL.Path)
	if err := h.service.Delete(id); err != nil {
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

/*
Для любой операции, объявляемой объектом, должны быть заданы: имя
операции, объекты, передаваемые в качестве параметров, и значение, возвращаемое операцией. Эту триаду называют сигнатурой операции. Множество
сигнатур всех определенных для объекта операций называется интерфейсом
этого объекта. Интерфейс описывает все множество запросов, которые можно отправить объекту. Любой запрос, сигнатура которого входит винтерфейс
объекта, может быть ему отправлен

Ассоциирование запроса с объектом и одной из
его операций во время выполнения называется динамическим связыванием.
*/
