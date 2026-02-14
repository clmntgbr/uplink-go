package project

import (
    "context"
    "uplink-go/domain"
)

func (s *Service) FindAll(ctx context.Context) ([]domain.Project, error) {
    return s.repo.FindAll(ctx)
}
