package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
)

var (
	addContactRe = regexp.MustCompile(`^/addcontact/$`)
)

type AddContactHandler struct {
	service addContactService
}

type addContactService interface {
	AddContact(ctx context.Context, login, addLogin string) error
}

type AddContactRequest struct {
	Login    string `json:"login"`
	AddLogin string `json:"addLogin"`
}

func NewAddContactHanlder(s addContactService) *AddContactHandler {
	return &AddContactHandler{
		service: s,
	}
}

func (h *AddContactHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !protectedHandler(w, r) {
		return
	}

	switch {
	case r.Method == http.MethodPost && addContactRe.MatchString(r.URL.Path):
		h.AddContact(w, r)
	}
}

func (h *AddContactHandler) AddContact(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), duringResponse)
	defer cancel()

	var req AddContactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}

	if err := h.service.AddContact(ctx, req.Login, req.AddLogin); err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
