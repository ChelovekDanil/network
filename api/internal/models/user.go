package models

type User struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	PassHash string `json:"passhash"`
}

type Token struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}
