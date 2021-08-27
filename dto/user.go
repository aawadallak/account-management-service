package dto

type User struct {
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Code     string `json:"code,omitempty"`
}
