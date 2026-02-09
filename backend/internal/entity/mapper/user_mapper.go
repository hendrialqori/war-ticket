package mapper

import (
	"github.com/hendrialqori/war-ticket/backend/internal/entity"
	"github.com/hendrialqori/war-ticket/backend/internal/model"
)

// mapper from model to entity
func ToUserEntity(m *model.UserModel) *entity.User {
	if m == nil {
		return nil
	}

	var act *entity.UserActivation

	// null check refers to scanning property on repository
	if m.Activation != nil {
		act = &entity.UserActivation{
			ID:        m.Activation.ID,
			Email:     m.Activation.Email,
			IsActive:  m.Activation.IsActive,
			CreatedAt: m.Activation.CreatedAt,
			UpdatedAt: m.Activation.UpdatedAt,
		}
	}

	return &entity.User{
		ID:           m.ID,
		Email:        m.Email,
		Username:     m.Username,
		Role:         m.Role,
		HashPassword: m.HashPassword,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
		Activation:   act,
	}
}
