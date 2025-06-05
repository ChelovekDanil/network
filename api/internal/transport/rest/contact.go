package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
)

var (
	addContactRe      = regexp.MustCompile(`^/contact/add/$`)
	deleteContactRe   = regexp.MustCompile(`^/contact/delete/$`)
	getContactsRe     = regexp.MustCompile(`^/contact/.*/$`)
	messageContactRe  = regexp.MustCompile(`^/contact/message/$`)
	getMessageContact = regexp.MustCompile(`^/contact/getmessage/$`)
)

type ContactHandler struct {
	service contactService
}

type contactService interface {
	AddContact(ctx context.Context, login, addLogin string) error
	DeleteContact(ctx context.Context, login, deleteLogin string) error
	GetAll(ctx context.Context, login string) ([]string, error)
	Message(ctx context.Context, firstLogin, lastLogin, message string) error
	GetMessage(ctx context.Context, firstLogin, lastLogin string) ([][]string, error)
}

type ContactRequest struct {
	FirstLogin string `json:"firstLogin"`
	LastLogin  string `json:"lastLogin"`
}

type MessageRequest struct {
	FirstLogin string `json:"firstLogin"`
	LastLogin  string `json:"lastLogin"`
	Message    string `json:"message"`
}

type GetMessageResponse struct {
	First []string `json:"first"`
	Last  []string `json:"last"`
}

func NewContactHanlder(s contactService) *ContactHandler {
	return &ContactHandler{
		service: s,
	}
}

func (h *ContactHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !protectedHandler(w, r) {
		return
	}

	switch {
	case r.Method == http.MethodPost && addContactRe.MatchString(r.URL.Path):
		h.AddContact(w, r)
	case r.Method == http.MethodPost && deleteContactRe.MatchString(r.URL.Path):
		h.DeleteContact(w, r)
	case r.Method == http.MethodGet && getContactsRe.MatchString(r.URL.Path):
		h.GetContacts(w, r)
	case r.Method == http.MethodPost && messageContactRe.MatchString(r.URL.Path):
		h.Message(w, r)
	case r.Method == http.MethodPost && getMessageContact.MatchString(r.URL.Path):
		h.GetMessage(w, r)
	}
}

func (h *ContactHandler) AddContact(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), duringResponse)
	defer cancel()

	var req ContactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}

	if err := h.service.AddContact(ctx, req.FirstLogin, req.LastLogin); err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ContactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), duringResponse)
	defer cancel()

	var req ContactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}

	if err := h.service.DeleteContact(ctx, req.FirstLogin, req.LastLogin); err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ContactHandler) GetContacts(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), duringResponse)
	defer cancel()

	login := getParamFromPath(r.URL.Path)

	users, err := h.service.GetAll(ctx, login)
	if err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}

	jsonUsers, err := json.Marshal(users)
	if err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}
	writeResponseOKWithData(w, jsonUsers)
}

func (h *ContactHandler) Message(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), duringResponse)
	defer cancel()

	var req MessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}

	if err := h.service.Message(ctx, req.FirstLogin, req.LastLogin, req.Message); err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ContactHandler) GetMessage(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), duringResponse)
	defer cancel()

	var req ContactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}

	messages, err := h.service.GetMessage(ctx, req.FirstLogin, req.LastLogin)
	if err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}

	response := GetMessageResponse{First: messages[0], Last: messages[1]}

	jsonUsers, err := json.Marshal(response)
	if err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}
	writeResponseOKWithData(w, jsonUsers)
}
