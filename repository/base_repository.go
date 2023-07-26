package repository

type BaseRepository[T any] interface {
	Create(payload T) error
	List() ([]T, error)
	GetId(id string) (T, error)
	Update(payload T) error
	Delete(id string) error
}
