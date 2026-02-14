package dto

type HydraResponse[T any] struct {
	CurrentPage  int `json:"current_page"`
	ItemsPerPage int `json:"items_per_page"`
	TotalPages   int `json:"total_pages"`
	TotalItems   int `json:"total_items"`
	Member       []T `json:"member"`
}

func NewHydraResponse[T any](items []T, current_page int, items_per_page int, total_items int) *HydraResponse[T] {
	total_pages := (total_items + items_per_page - 1) / items_per_page
	if total_pages == 0 {
		total_pages = 1
	}

	return &HydraResponse[T]{
		CurrentPage:  current_page,
		ItemsPerPage: items_per_page,
		TotalPages:   total_pages,
		TotalItems:   total_items,
		Member:       items,
	}
}
