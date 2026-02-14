package project

import (
	"uplink-go/service"
)

type ProjectHandler struct {
	getProjectsService *service.GetProjectsService
	createProjectService *service.CreateProjectService
}

func NewProjectHandler(getProjectsService *service.GetProjectsService, createProjectService *service.CreateProjectService) *ProjectHandler {
	return &ProjectHandler{
		getProjectsService: getProjectsService,
		createProjectService: createProjectService,
	}
}

