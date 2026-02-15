package project

import (
	"context"
	"uplink-go/ctxutil"
	"uplink-go/dto"
)

func (s *Service) FindAll(ctx context.Context) (*dto.HydraResponse[dto.ProjectResponse], error) {
	if userID, ok := ctxutil.GetUserIDFromContext(ctx); ok {
		if activeProjectID, err := s.repo.FindActiveProject(ctx, userID); err == nil {
			ctx = ctxutil.WithActiveProjectID(ctx, activeProjectID)
		}
	}

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
