package entity

import "mime/multipart"

type MissionRepositoryInterface interface {
	Create(data Mission) error
}

type MissionServiceInterface interface {
	CreateMission(image *multipart.FileHeader, data Mission) error
}
