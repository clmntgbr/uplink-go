package repository

import (
	"context"
	"errors"
	"uplink-go/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Create(ctx context.Context, project *domain.Project) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("Users").Create(project).Error; err != nil {
			return err
		}

		if len(project.Users) > 0 {
			if err := tx.Model(project).Association("Users").Append(project.Users); err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *ProjectRepository) FindAll(ctx context.Context, userID uuid.UUID) ([]domain.Project, error) {
	var user domain.User
	user.ID = userID

	var projects []domain.Project
	err := r.db.WithContext(ctx).Model(&user).Association("Projects").Find(&projects)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *ProjectRepository) FindByID(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*domain.Project, error) {
	var project domain.Project
	err := r.db.WithContext(ctx).
		Joins("JOIN user_projects ON user_projects.project_id = projects.id").
		Where("projects.id = ? AND user_projects.user_id = ?", id, userID).
		First(&project).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *ProjectRepository) Delete(ctx context.Context, id uuid.UUID, userID uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Project{}, "id = ? AND user_id = ?", id, userID).Error
}

func (r *ProjectRepository) FindActiveProject(ctx context.Context, userID uuid.UUID) (*uuid.UUID, error) {
	var result struct {
		ActiveProjectID *uuid.UUID `gorm:"column:active_project_id"`
	}
	
	err := r.db.WithContext(ctx).
		Model(&domain.User{}).
		Select("active_project_id").
		Where("id = ?", userID).
		First(&result).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	
	return result.ActiveProjectID, nil
}

