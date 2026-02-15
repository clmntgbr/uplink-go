package project

import (
	"context"
	"uplink-go/dto"

	"github.com/google/uuid"
)

func (s *Service) FindAll(ctx context.Context, userID uuid.UUID) (*dto.HydraResponse[dto.ProjectResponse], error) {
	projects, err := s.repo.FindAll(ctx, userID)
	if err != nil {
		return nil, err
	}

	activeProjectID, err := s.repo.FindActiveProject(ctx, userID)
	if err != nil {
		return nil, err
	}

	return dto.NewHydraResponse(
		dto.ToProjectsResponse(projects, activeProjectID),
		1,
		10,
		len(projects),
	), nil
}
