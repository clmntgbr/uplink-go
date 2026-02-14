package project

import (
    "context"
    "errors"
	"uplink-go/domain"
)

type CreateInput struct {
    Name  string
}

func (s *Service) Create(ctx context.Context, input CreateInput) (*domain.Project, error) {
    if input.Name == "" {
        return nil, errors.New("name required")
    }

    project := &domain.Project{
        Name: input.Name,
    }

    if err := s.repo.Create(ctx, project); err != nil {
        return nil, err
    }

    return project, nil
}
