package dto

type UserUpdateEmail struct {
	Oldemail string `json:"oldEmail,omitempty"`
	Code     string `json:"code,omitempty"`
	NewEmail string `json:"newEmail,omitempty"`
}
