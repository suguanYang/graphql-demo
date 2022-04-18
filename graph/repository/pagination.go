package repository

const (
	LIMIT  = 10
	OFFSET = 0
)

type Pagination struct {
	Offset int
	Limit  int
}

func NewPagination(limit *int, offset *int) *Pagination {
	pagination := &Pagination{
		Offset: OFFSET,
		Limit:  LIMIT,
	}
	if offset != nil {
		pagination.Offset = *offset
	}
	if limit != nil {
		pagination.Limit = *limit
	}

	return pagination
}
