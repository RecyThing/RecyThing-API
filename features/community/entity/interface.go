package entity

import (
	"mime/multipart"
	"recything/utils/pagination"
)

type CommunityRepositoryInterface interface {
	CreateCommunity(image *multipart.FileHeader, data CommunityCore) error
	GetAllCommunity(page, limit int, search string) ([]CommunityCore, pagination.PageInfo, int, error)
	GetCommunityById(id string) (CommunityCore, error)
	GetByName(name string) (CommunityCore, error)
	UpdateCommunityById(id string, image *multipart.FileHeader, data CommunityCore) error
	DeleteCommunityById(id string) error
}

type CommunityServiceInterface interface {
	CreateCommunity(image *multipart.FileHeader, data CommunityCore) error
	GetAllCommunity(page, limit, search string) ([]CommunityCore, pagination.PageInfo, int, error)
	GetCommunityById(id string) (CommunityCore, error)
	UpdateCommunityById(id string, image *multipart.FileHeader, data CommunityCore) error
	DeleteCommunityById(id string) error
}
