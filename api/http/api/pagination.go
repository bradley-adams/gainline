package api

type PaginationRequest struct {
	Page     int `form:"page" validate:"omitempty,gte=1"`
	PageSize int `form:"page_size" validate:"omitempty,gte=1,lte=100"`
}

type PaginationMeta struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

type PaginatedResponse[T any] struct {
	Data       []T            `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}

func (q *PaginationRequest) SetDefaults() {
	if q.Page == 0 {
		q.Page = 1
	}
	if q.PageSize == 0 {
		q.PageSize = 20
	}
}

func (q *PaginationRequest) Offset() int {
	return (q.Page - 1) * q.PageSize
}
