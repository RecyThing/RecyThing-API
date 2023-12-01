package entity

import "recything/utils/pagination"

type TrashExchangeRepositoryInterface interface {
	CreateTrashExchange(data TrashExchangeCore) (TrashExchangeCore, error)
	CreateTrashExchangeDetails(data TrashExchangeDetailCore) (TrashExchangeDetailCore, error)
	GetTrashExchangeById(id string) (TrashExchangeCore, error)
	GetAllTrashExchange(page, limit int, search string) ([]TrashExchangeCore, pagination.PageInfo, error)
	DeleteTrashExchangeById(id string) error
}

type TrashExchangeServiceInterface interface {
	CreateTrashExchange(data TrashExchangeCore) (TrashExchangeCore, error)
	GetTrashExchangeById(id string) (TrashExchangeCore, error)
	GetAllTrashExchange(page, limit int, name, address string) ([]TrashExchangeCore, pagination.PageInfo, error)
	DeleteTrashExchangeById(id string) error
}
