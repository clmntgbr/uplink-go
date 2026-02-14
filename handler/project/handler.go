package project

import (
	"uplink-go/service/project"
)

type ProjectHandler struct {
	projectService *project.Service
}

func NewProjectHandler(service *project.Service) *ProjectHandler {
	return &ProjectHandler{
		projectService: service,
	}
}

