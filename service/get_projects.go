package service

import (
	"uplink-go/dto"
	"uplink-go/repository"

	"github.com/google/uuid"
)

type GetProjectsService struct {
	userRepo *repository.UserRepository
}

func NewGetProjectsService(userRepo *repository.UserRepository) *GetProjectsService {
	return &GetProjectsService{
		userRepo: userRepo,
	}
}

func (s *GetProjectsService) Projects(userID uuid.UUID) (dto.HydraResponse[dto.ProjectResponse], error) {
	projects, err := s.userRepo.FindProjectsByUserID(userID)
	if err != nil {
		return dto.HydraResponse[dto.ProjectResponse]{}, err
	}

	return dto.NewHydraResponse(
		dto.ToProjectsResponse(projects),
		1,
		10,
		len(projects),
	), nil
}