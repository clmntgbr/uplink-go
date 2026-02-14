package repository

import (
	"context"
	"uplink-go/domain"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Create(ctx context.Context, project *domain.Project) error {
	return r.db.WithContext(ctx).Create(project).Error
}

func (r *ProjectRepository) FindAll(ctx context.Context) ([]domain.Project, error) {
	var projects []domain.Project
	err := r.db.WithContext(ctx).Find(&projects).Error
	return projects, err
}

func (r *ProjectRepository) FindByID(ctx context.Context, id string) (*domain.Project, error) {
	var project domain.Project
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&project).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *ProjectRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&domain.Project{}, "id = ?", id).Error
}