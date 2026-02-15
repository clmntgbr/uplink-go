package project

import (
	"context"
	"errors"
	"uplink-go/ctxutil"
	"uplink-go/dto"
	apperrors "uplink-go/errors"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

func (s *Service) FindById(ctx context.Context, projectID uuid.UUID) (dto.ProjectResponse, error) {
	if userID, ok := ctxutil.GetUserIDFromContext(ctx); ok {
		if activeProjectID, err := s.repo.FindActiveProject(ctx, userID); err == nil {
			ctx = ctxutil.WithActiveProjectID(ctx, activeProjectID)
		}
	}

	project, err := s.repo.FindByID(ctx, projectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.ProjectResponse{}, apperrors.ErrProjectNotFound
		}
		return dto.ProjectResponse{}, err
	}

	return dto.ToProjectResponse(*project), nil
}
