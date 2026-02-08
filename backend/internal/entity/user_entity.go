package entity

import "time"

type User struct {
	ID           string
	Email        string
	Username     string
	HashPassword string
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Activation *UserActivation
}

type UserActivation struct {
	ID        string
	UserID    string
	Email     string
	IsActive  int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
