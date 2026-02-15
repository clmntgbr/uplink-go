package dto

import (
	"time"
	"uplink-go/domain"

	"github.com/google/uuid"
)

type ProjectResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsActive  bool      `json:"isActive"`
}

type CreateInput struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}

type ActivateInput struct {
	ProjectID uuid.UUID `json:"projectId" validate:"required,uuid"`
}

func ToProjectResponse(p domain.Project) ProjectResponse {
	return ProjectResponse{
		ID:   p.ID,
		Name: p.Name,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		IsActive: p.IsActive,
	}
}

func ToProjectsResponse(projects []domain.Project) []ProjectResponse {
	result := make([]ProjectResponse, len(projects))
	for i, p := range projects {
		result[i] = ToProjectResponse(p)
	}
	return result
}