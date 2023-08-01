package dto

type PaginationParam struct {
	Page   int
	Offset int
	Limit  int
}

type PaginationQuery struct {
	Page int
	Take int
	Skip int
}

type Paging struct {
	Page        int `json:"page"`
	RowsPerPage int `json:"rows_per_page"`
	TotalRows   int `json:"total_rows"`
	TotalPages  int `json:"total_pages"`
}
