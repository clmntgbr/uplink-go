package project

import (
	"context"
	"uplink-go/dto"

	"github.com/google/uuid"
)

func (s *Service) FindAll(ctx context.Context, userID uuid.UUID) ([]dto.ProjectResponse, error) {
    projects, err := s.repo.FindAll(ctx, userID)
	if err != nil {
		return nil, err
	}
	
	resp := dto.ToProjectsResponse(projects)
	return resp, nil
}
