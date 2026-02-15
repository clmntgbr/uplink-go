package repository

import (
	"context"
	"errors"
	"uplink-go/domain"
	"uplink-go/ctxutil"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Create(ctx context.Context, project *domain.Project, userID uuid.UUID) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("Users").Create(project).Error; err != nil {
			return err
		}

		var user domain.User
		user.ID = userID
		if err := tx.Model(project).Association("Users").Append(&user); err != nil {
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

func (r *ProjectRepository) FindAll(ctx context.Context) ([]domain.Project, error) {
	userID, ok := ctxutil.GetUserIDFromContext(ctx)
	if !ok {
		return nil, errors.New("user ID not found in context")
	}

	var projects []domain.Project

	err := r.db.WithContext(ctx).
		Joins("JOIN user_projects ON user_projects.project_id = projects.id").
		Where("user_projects.user_id = ?", userID).
		Find(&projects).Error

	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *ProjectRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Project, error) {
	userID, ok := ctxutil.GetUserIDFromContext(ctx)
	if !ok {
		return nil, errors.New("user ID not found in context")
	}

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

func (r *ProjectRepository) Delete(ctx context.Context, id uuid.UUID) error {
	userID, ok := ctxutil.GetUserIDFromContext(ctx)
	if !ok {
		return errors.New("user ID not found in context")
	}

	var project domain.Project
	err := r.db.WithContext(ctx).
		Joins("JOIN user_projects ON user_projects.project_id = projects.id").
		Where("projects.id = ? AND user_projects.user_id = ?", id, userID).
		First(&project).Error

	if err != nil {
		return err
	}

	return r.db.WithContext(ctx).Delete(&project).Error
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

func (r *ProjectRepository) ActivateProject(ctx context.Context, projectID uuid.UUID) error {
	userID, ok := ctxutil.GetUserIDFromContext(ctx)
	if !ok {
		return errors.New("user ID not found in context")
	}

	var project domain.Project
	err := r.db.WithContext(ctx).
		Joins("JOIN user_projects ON user_projects.project_id = projects.id").
		Where("projects.id = ? AND user_projects.user_id = ?", projectID, userID).
		First(&project).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("project not found or access denied")
		}
		return err
	}

	return r.db.WithContext(ctx).
		Model(&domain.User{}).
		Where("id = ?", userID).
		Update("active_project_id", projectID).Error
}