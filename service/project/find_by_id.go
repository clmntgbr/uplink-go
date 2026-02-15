package project

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"uplink-go/dto"
	apperrors "uplink-go/errors"

	"github.com/google/uuid"
)

func (s *Service) FindById(ctx context.Context, userID uuid.UUID, projectID uuid.UUID) (dto.ProjectResponse, error) {
	project, err := s.repo.FindByID(ctx, projectID, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.ProjectResponse{}, apperrors.ErrProjectNotFound
		}
		return dto.ProjectResponse{}, err
	}

	activeProjectID, err := s.repo.FindActiveProject(ctx, userID)
	if err != nil {
		return dto.ProjectResponse{}, err
	}

	return dto.ToProjectResponse(*project, activeProjectID), nil
}
