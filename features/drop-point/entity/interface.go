package entity

import "recything/utils/pagination"

type DropPointRepositoryInterface interface {
	CreateDropPoint(data DropPointCore) (DropPointCore, error)
	UpdateDropPointById(id string, data DropPointCore) (DropPointCore, error)
	GetDropPointById(id string) (DropPointCore, error)
	GetAllDropPoint(page, limit int, name, address string) ([]DropPointCore, pagination.PageInfo, error)
	DeleteDropPointById(id string) error
}

type DropPointServiceInterface interface {
	CreateDropPoint(data DropPointCore) (DropPointCore, error)
	UpdateDropPointById(id string, data DropPointCore) (DropPointCore, error)
	GetDropPointById(id string) (DropPointCore, error)
	GetAllDropPoint(page, limit int, name, address string) ([]DropPointCore, pagination.PageInfo, error)
	DeleteDropPointById(id string) error
}
