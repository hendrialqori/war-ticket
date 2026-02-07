package dto

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,excludes= "`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type SetActiveRequest struct {
	Email string `json:"email" validate:"required,email"`
}
