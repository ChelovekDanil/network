package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/ChelovekDanil/network/internal/models"
)

var (
	authLoginRe   = regexp.MustCompile(`^/auth/login/$`)
	authRefreshRe = regexp.MustCompile(`^/auth/refresh/$`)
)

type AuthHandler struct {
	service authService
}

type tokens struct {
	Access  string `json:"access-token"`
	Refresh string `json:"refresh-token"`
}

type requestReLogin struct {
	models.User
	RefreshToken string `json:"refresh-token"`
}

type authService interface {
	Login(ctx context.Context, user models.User) ([]string, error)
	ReLogin(ctx context.Context, user models.User, refreshToken string) ([]string, error)
}

func NewAuthHandler(s authService) *AuthHandler {
	return &AuthHandler{
		service: s,
	}
}

func (h *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && authLoginRe.MatchString(r.URL.Path):
		h.Login(w, r)
	case r.Method == http.MethodPost && authRefreshRe.MatchString(r.URL.Path):
		h.ReLogin(w, r)
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), duringResponse)
	defer cancel()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}

	t, err := h.service.Login(ctx, user)
	if err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(tokens{t[0], t[1]})
}

func (h *AuthHandler) ReLogin(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), duringResponse)
	defer cancel()

	var relog requestReLogin
	if err := json.NewDecoder(r.Body).Decode(&relog); err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}

	t, err := h.service.ReLogin(ctx, relog.User, relog.RefreshToken)
	if err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(tokens{t[0], t[1]})
}
