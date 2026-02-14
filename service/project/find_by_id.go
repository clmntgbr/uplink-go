package project

import (
	"context"
	"errors"
	"uplink-go/dto"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

func (s *Service) FindById(ctx context.Context, userID uuid.UUID, projectID uuid.UUID) (dto.ProjectResponse, error) {
	project, err := s.repo.FindByID(ctx, projectID, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.ProjectResponse{}, errors.New("project not found")
		}
		return dto.ProjectResponse{}, err
	}

	return dto.ToProjectResponse(*project), nil
}
