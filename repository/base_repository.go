package repository

import "asetku-bukan-asetmu/model/dto"

type BaseRepository[T any] interface {
	Create(payload T) error
	List() ([]T, error)
	Get(id string) (T, error)
	Update(payload T) error
	Delete(id string) error
}

type BaseRepositoryPaging[T any] interface {
	Pagination(requestPaging dto.PaginationQueryParam, query ...string) ([]T, dto.PaginationResponse, error)
}
