package dto

import (
	"time"
	"uplink-go/domain"

	"github.com/google/uuid"
)

type ProjectResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToProjectResponse(p domain.Project) ProjectResponse {
	return ProjectResponse{
		ID:   p.ID,
		Name: p.Name,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func ToProjectsResponse(projects []domain.Project) []ProjectResponse {
	result := make([]ProjectResponse, len(projects))
	for i, p := range projects {
		result[i] = ToProjectResponse(p)
	}
	return result
}