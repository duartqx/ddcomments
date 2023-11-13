package http

import "time"

type LoginResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
	Status    string    `json:"status"`
}
