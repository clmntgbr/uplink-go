package dto

import (
	"uplink-go/domain"

	"github.com/google/uuid"
)

type EndpointResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
}

func ToEndpointResponse(p domain.Endpoint) EndpointResponse {
	return EndpointResponse{
		ID:   p.ID,
		Name: p.Name,
	}
}

func ToEndpointsResponse(endpoints []domain.Endpoint) []EndpointResponse {
	result := make([]EndpointResponse, len(endpoints))
	for i, p := range endpoints {
		result[i] = ToEndpointResponse(p)
	}
	return result
}