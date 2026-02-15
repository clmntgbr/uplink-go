package project

import (
	"context"
	"uplink-go/domain"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, project *domain.Project) error
	FindAll(ctx context.Context) ([]domain.Project, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Project, error)
	Delete(ctx context.Context, id uuid.UUID) error
	FindActiveProject(ctx context.Context, userID uuid.UUID) (*uuid.UUID, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
