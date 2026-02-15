package endpoint

import (
	"context"
	"uplink-go/dto"
)

func (s *Service) FindAll(ctx context.Context) (*dto.HydraResponse[dto.EndpointResponse], error) {
	endpoints, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return dto.NewHydraResponse(
		dto.ToEndpointsResponse(endpoints),
		1,
		10,
		len(endpoints),
	), nil
}
