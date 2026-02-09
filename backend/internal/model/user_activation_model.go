package model

import "time"

type UserActivationModel struct {
	ID       string `gorm:"column:id;primaryKey;type:char(36)"`
	UserID   string `gorm:"column:user_id;type:char(36);uniqueIndex"`
	Email    string `gorm:"column:email;type:varchar(100)"`
	IsActive int32  `gorm:"column:is_active"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`

	User *UserModel
}

func (UserActivationModel) TableName() string {
	return "user_activations"
}
