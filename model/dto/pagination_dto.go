package dto

type PaginationQueryParam struct {
	Page, Offset, Limit int
}

type PaginationReturn struct {
	CurrentPage, LimitRows, StartIndex int
}

type PaginationResponse struct {
	Page, RowsPerPage, TotalRows, TotalPages int
}
