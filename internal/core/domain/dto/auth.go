package dto

import "time"

type AuthRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthRegisterResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Role      string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type AuthLoginResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"status"`
	Token string `json:"token"`
}
