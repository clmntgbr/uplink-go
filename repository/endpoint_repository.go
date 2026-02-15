package repository

import (
	"context"
	"errors"
	"uplink-go/ctxutil"
	"uplink-go/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EndpointRepository struct {
	db *gorm.DB
}

func NewEndpointRepository(db *gorm.DB) *EndpointRepository {
	return &EndpointRepository{db: db}
}

func (r *EndpointRepository) FindAll(ctx context.Context) ([]domain.Endpoint, error) {
	userID, ok := ctxutil.GetUserIDFromContext(ctx)
	if !ok {
		return nil, errors.New("user ID not found in context")
	}

	activeProjectID, ok := ctxutil.GetActiveProjectID(ctx)
	if !ok || activeProjectID == nil {
		return nil, errors.New("no active project")
	}

	var endpoints []domain.Endpoint

	err := r.db.WithContext(ctx).
		Joins("JOIN projects ON projects.id = endpoints.project_id").
		Joins("JOIN user_projects ON user_projects.project_id = projects.id").
		Where("endpoints.project_id = ? AND user_projects.user_id = ?", *activeProjectID, userID).
		Find(&endpoints).Error

	if err != nil {
		return nil, err
	}
	return endpoints, nil
}

func (r *EndpointRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Endpoint, error) {
	userID, ok := ctxutil.GetUserIDFromContext(ctx)
	if !ok {
		return nil, errors.New("user ID not found in context")
	}

	activeProjectID, ok := ctxutil.GetActiveProjectID(ctx)
	if !ok || activeProjectID == nil {
		return nil, errors.New("no active project")
	}

	var endpoint domain.Endpoint

	err := r.db.WithContext(ctx).
		Joins("JOIN projects ON projects.id = endpoints.project_id").
		Joins("JOIN user_projects ON user_projects.project_id = projects.id").
		Where("endpoints.id = ? AND endpoints.project_id = ? AND user_projects.user_id = ?", id, *activeProjectID, userID).
		First(&endpoint).Error

	if err != nil {
		return nil, err
	}
	return &endpoint, nil
}
