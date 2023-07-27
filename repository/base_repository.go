package repository

import "github.com/albar2305/enigma-laundry-apps/model/dto"

type BaseRepository[T any] interface {
	Create(payload T) error
	List() ([]T, error)
	GetById(id string) (T, error)
	Update(payload T) error
	Delete(id string) error
}

type BaseRepositoryPaging[T any] interface {
	Paging(requestPaging dto.PaginationParam) ([]T, dto.Paging, error)
}
