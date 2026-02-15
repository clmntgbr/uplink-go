package project

import (
	"context"
	"uplink-go/dto"
)

func (s *Service) FindAll(ctx context.Context) (*dto.HydraResponse[dto.ProjectResponse], error) {
	projects, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return dto.NewHydraResponse(
		dto.ToProjectsResponse(projects),
		1,
		10,
		len(projects),
	), nil
}
