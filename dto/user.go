package dto

import (
	"time"
	"uplink-go/domain"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToUserResponse(p *domain.User) UserResponse {
	return UserResponse{
		ID:   p.ID,
		Email: p.Email,
		FirstName: p.FirstName,
		LastName: p.LastName,
		Avatar: p.Avatar,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}