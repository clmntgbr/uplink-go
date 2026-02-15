package endpoint

import "uplink-go/service/endpoint"

type EndpointHandler struct {
	endpointService *endpoint.Service
}

func NewEndpointHandler(service *endpoint.Service) *EndpointHandler {
	return &EndpointHandler{
		endpointService: service,
	}
}

