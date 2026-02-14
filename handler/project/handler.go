package project

import "uplink-go/repository"

type ProjectHandler struct {
	projectRepo *repository.ProjectRepository
	userRepo *repository.UserRepository
}

func NewProjectHandler(projectRepo *repository.ProjectRepository, userRepo *repository.UserRepository) *ProjectHandler {
	return &ProjectHandler{
		projectRepo: projectRepo,
		userRepo: userRepo,
	}
}

