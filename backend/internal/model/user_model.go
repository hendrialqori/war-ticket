package model

import "time"

type UserModel struct {
	ID           string    `gorm:"column:id;primaryKey;type:char(36)"`
	Email        string    `gorm:"column:email;uniqueIndex;type:char(36)"`
	Username     string    `gorm:"column:username;uniqueIndex;type:varchar(100)"`
	HashPassword string    `gorm:"column:hash_password;type:varchar(255)"`
	Role         string    `gorm:"column:role;type:varchar(16)"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`

	Activation *UserActivationModel `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (UserModel) TableName() string {
	return "users" // naming table
}
