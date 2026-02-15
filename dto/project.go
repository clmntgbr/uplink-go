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

func ToProjectResponse(p domain.Project, activeProjectID *uuid.UUID) ProjectResponse {
	isActive := false
	if activeProjectID != nil && *activeProjectID == p.ID {
		isActive = true
	}

	return ProjectResponse{
		ID:   p.ID,
		Name: p.Name,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		IsActive:  isActive,
	}
}

func ToProjectsResponse(projects []domain.Project, activeProjectID *uuid.UUID) []ProjectResponse {
	result := make([]ProjectResponse, len(projects))
	for i, p := range projects {
		result[i] = ToProjectResponse(p, activeProjectID)
	}
	return result
}