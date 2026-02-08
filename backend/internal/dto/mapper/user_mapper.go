package mapper

import (
	"github.com/hendrialqori/war-ticket/backend/internal/dto"
	"github.com/hendrialqori/war-ticket/backend/internal/entity"
)

func ToUserDTO(u *entity.User) *dto.UserResponse {
	if u == nil {
		return nil
	}

	out := &dto.UserResponse{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}

	if u.Activation != nil {
		out.Activation = &dto.UserActivationResponse{
			IsActive:   u.Activation.IsActive == 1,
			ActivateAt: u.Activation.CreatedAt,
		}
	}
	return out
}
