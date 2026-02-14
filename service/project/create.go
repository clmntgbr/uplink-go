package project

import (
	"context"
	"errors"
	"uplink-go/domain"
	"uplink-go/dto"

	"github.com/google/uuid"
)

type CreateInput struct {
    Name  string
}

func (s *Service) Create(ctx context.Context, input CreateInput, userID uuid.UUID) (*dto.ProjectResponse, error) {
    if input.Name == "" {
        return nil, errors.New("name required")
    }

    project := &domain.Project{
        Name: input.Name,
				Users: []domain.User{
					{
						ID: userID,
					},
				},
    }

    if err := s.repo.Create(ctx, project); err != nil {
        return nil, err
    }

    resp := dto.ToProjectResponse(*project)
		return &resp, nil
}
