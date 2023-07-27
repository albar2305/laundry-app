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
	Page        int
	RowsPerPage int
	TotalRows   int
	TotalPages  int
}
