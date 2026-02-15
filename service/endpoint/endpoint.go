package endpoint

import (
	"context"
	"uplink-go/domain"

	"github.com/google/uuid"
)

type Repository interface {
	FindAll(ctx context.Context) ([]domain.Endpoint, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Endpoint, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
