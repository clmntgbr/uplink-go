package project

import (
	"context"
	"uplink-go/domain"
	"uplink-go/dto"

	"github.com/google/uuid"
)

func (s *Service) Create(ctx context.Context, input dto.CreateInput, userID uuid.UUID) (*dto.ProjectResponse, error) {
	project := &domain.Project{
		Name: input.Name,
		Users: []domain.User{
			{
				ID: userID,
			},
		},
	}

	if err := s.repo.Create(ctx, project, userID); err != nil {
		return nil, err
	}

	resp := dto.ToProjectResponse(*project)
	return &resp, nil
}
