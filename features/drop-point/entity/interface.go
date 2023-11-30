package entity

import "recything/utils/pagination"

type DropPointRepositoryInterface interface {
	CreateDropPoint(data DropPointsCore) error
	GetAllDropPoint(page, limit int, search string) ([]DropPointsCore, pagination.PageInfo, error)
	GetDropPointById(id string) (DropPointsCore, error)
	UpdateDropPointById(id string, data DropPointsCore) error
	DeleteDropPointById(id string) error
}

type DropPointServiceInterface interface {
	CreateDropPoint(data DropPointsCore) error
	GetAllDropPoint(page, limit int, search string) ([]DropPointsCore, pagination.PageInfo, error)
	GetDropPointById(id string) (DropPointsCore, error)
	UpdateDropPointById(id string, data DropPointsCore) error
	DeleteDropPointById(id string) error
}
