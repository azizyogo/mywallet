package pagination

type PaginationParams struct {
	Page  int
	Limit int
}

func NewPaginationParams(page, limit int) PaginationParams {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	return PaginationParams{Page: page, Limit: limit}
}

func (p PaginationParams) Offset() int {
	return (p.Page - 1) * p.Limit
}

func CalculateTotalPages(total int64, limit int) int {
	if limit == 0 {
		return 0
	}
	pages := int(total) / limit
	if int(total)%limit > 0 {
		pages++
	}
	return pages
}
