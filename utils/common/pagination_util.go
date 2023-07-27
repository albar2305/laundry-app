package common

import (
	"math"
	"os"
	"strconv"

	"github.com/albar2305/enigma-laundry-apps/model/dto"
	"github.com/albar2305/enigma-laundry-apps/utils/exceptions"
)

func GetPaginationParams(param dto.PaginationParam) dto.PaginationQuery {
	err := LoadEnv()
	exceptions.CheckErr(err)
	var (
		page, take, skip int
	)

	if param.Page > 0 {
		page = param.Page
	} else {
		page = 1
	}

	if param.Limit == 0 {
		n, _ := strconv.Atoi(os.Getenv("DEFAULT_ROW_PER_PAGE"))
		take = n
	} else {
		take = param.Limit
	}

	if page > 0 {
		skip = (page - 1) * take
	} else {
		skip = 0
	}

	return dto.PaginationQuery{
		Page: page,
		Take: take,
		Skip: skip,
	}
}

func Pagination(page, limit, totalRows int) dto.Paging {
	return dto.Paging{
		Page:        page,
		RowsPerPage: limit,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(limit))),
	}
}
