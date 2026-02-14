package project

import (
    "context"
    "uplink-go/domain"
)

type Repository interface {
    Create(ctx context.Context, project *domain.Project) error
    FindAll(ctx context.Context) ([]domain.Project, error)
    FindByID(ctx context.Context, id string) (*domain.Project, error)
    Delete(ctx context.Context, id string) error
}

type Service struct {
    repo Repository
}

func New(repo Repository) *Service {
    return &Service{
        repo: repo,
    }
}
