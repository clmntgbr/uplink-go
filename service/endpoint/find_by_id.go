package endpoint

import (
	"context"
	"errors"
	"uplink-go/dto"
	apperrors "uplink-go/errors"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

func (s *Service) FindById(ctx context.Context, endpointID uuid.UUID) (dto.EndpointResponse, error) {
	endpoint, err := s.repo.FindByID(ctx, endpointID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.EndpointResponse{}, apperrors.ErrEndpointNotFound
		}
		return dto.EndpointResponse{}, err
	}

	return dto.ToEndpointResponse(*endpoint), nil
}
