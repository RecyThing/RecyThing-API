package entity

import (
	"mime/multipart"
	"recything/utils/pagination"
)

type MissionRepositoryInterface interface {
	Create(input Mission)error
	FindAll(page, limit int, filter string) ([]Mission, pagination.PageInfo, int, error)
	GetCount(filter string) (int, error)
}

type MissionServiceInterface interface {
	CreateMission(image *multipart.FileHeader, data Mission) error
	FindAll(page, limit string, filter string) ([]Mission, pagination.PageInfo, int, error)
}
