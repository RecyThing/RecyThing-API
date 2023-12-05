package entity

import "recything/utils/pagination"

type RecybotRepositoryInterface interface {
	Create(recybot RecybotCore) (RecybotCore, error)
	Update(idData string, data RecybotCore) (RecybotCore, error)
	Delete(idData string) error
	FindAll(page, limit int, filter, search string) ([]RecybotCore, pagination.PageInfo, int, error)
	GetAll() ([]RecybotCore, error)
	GetById(idData string) (RecybotCore, error)
	GetCount(filter, search string) (int, error)
}

type RecybotServiceInterface interface {
	CreateData(recybot RecybotCore) (RecybotCore, error)
	UpdateData(idData string, data RecybotCore) (RecybotCore, error)
	DeleteData(idData string) error
	FindAllData(filter, search, page, limit string) ([]RecybotCore, pagination.PageInfo, int, error)
	GetById(idData string) (RecybotCore, error)
	GetPrompt(question string) (string, error)
}
