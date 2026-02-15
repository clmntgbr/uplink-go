package project

import (
	"context"
	"uplink-go/dto"
)

func (s *Service) ActivateProject(ctx context.Context, input dto.ActivateInput) error {
	return s.repo.ActivateProject(ctx, input.ProjectID)
}
