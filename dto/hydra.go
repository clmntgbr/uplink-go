package dto

type HydraResponse[T any] struct {
	CurrentPage  int `json:"currentPage"`
	ItemsPerPage int `json:"itemsPerPage"`
	TotalPages   int `json:"totalPages"`
	TotalItems   int `json:"totalItems"`
	Member       []T `json:"member"`
}

func NewHydraResponse[T any](items []T, currentPage int, itemsPerPage int, totalItems int) HydraResponse[T] {
	totalPages := (totalItems + itemsPerPage - 1) / itemsPerPage
	if totalPages == 0 {
		totalPages = 1
	}
	
	return HydraResponse[T]{
		CurrentPage:  currentPage,
		ItemsPerPage: itemsPerPage,
		TotalPages:   totalPages,
		TotalItems:   totalItems,
		Member:       items,
	}
}