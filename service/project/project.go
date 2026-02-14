package project

import (
	"context"
	"uplink-go/domain"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, project *domain.Project) error
	FindAll(ctx context.Context, userID uuid.UUID) ([]domain.Project, error)
	FindByID(ctx context.Context, id string, userID uuid.UUID) (*domain.Project, error)
	Delete(ctx context.Context, id string, userID uuid.UUID) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
