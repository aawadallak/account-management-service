package dto

type UserUpdatePassword struct {
	Email       string `json:"email,omitempty"`
	Username    string `json:"username,omitempty"`
	OldPassword string `json:"oldPassword,omitempty"`
	NewPassword string `json:"newPassword,omitempty"`
}
