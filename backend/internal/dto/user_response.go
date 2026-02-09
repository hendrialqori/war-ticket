package dto

import "time"

type UserResponse struct {
	ID         string                  `json:"id"`
	Email      string                  `json:"email"`
	Username   string                  `json:"username"`
	Role       string                  `json:"role"`
	CreatedAt  time.Time               `json:"created_at"`
	Activation *UserActivationResponse `json:"activation,omitempty"`
}

type UserActivationResponse struct {
	IsActive   bool      `json:"is_active"`
	ActivateAt time.Time `json:"activate_at"`
}
