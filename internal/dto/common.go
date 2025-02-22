package dto

type Pagination struct {
	Page  *int `json:"page" validate:"omitempty,min=1"`
	Limit *int `json:"limit" validate:"omitempty,min=1"`
}

type PaginationMeta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"totalPages"`
}
