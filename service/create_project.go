package service

import (
	"errors"

	"uplink-go/domain"
	"uplink-go/dto"
	"uplink-go/repository"

	"github.com/google/uuid"
)

type CreateProjectService struct {
	projectRepo *repository.ProjectRepository
}

type CreateProjectRequest struct {
	Name string `json:"name"`
}

func NewCreateProjectService(projectRepo *repository.ProjectRepository) *CreateProjectService {
	return &CreateProjectService{
		projectRepo: projectRepo,
	}
}

func (s *CreateProjectService) Create(userID uuid.UUID, req CreateProjectRequest) (*dto.ProjectResponse, error) {
	if req.Name == "" {
		return nil, errors.New("name is required")
	}

	project := &domain.Project{
		Name: req.Name,
		Users: []domain.User{
			{
				ID: userID,
			},
		},
	}

	if err := s.projectRepo.Create(project); err != nil {
		return nil, err
	}

	resp := dto.ToProjectResponse(*project)
	return &resp, nil
}