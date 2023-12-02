package entity

import (
	"mime/multipart"
	"recything/utils/pagination"
)

type MissionRepositoryInterface interface {
	CreateMission(input Mission) error
	FindAllMission(page, limit int, search, status string) ([]Mission, pagination.PageInfo, int, error)
	GetCount(status, search string) (int, error)
	SaveChangesStatusMission(data Mission) error
	UpdateMission(missionID string, data Mission) error
	UpdateMissionStage(missionStageID string, data MissionStage) error
	GetImageURL(missionID string)(string, error)
}

type MissionServiceInterface interface {
	CreateMission(image *multipart.FileHeader, data Mission) error
	FindAllMission(page, limit, search, status string) ([]Mission, pagination.PageInfo, int, error)
	UpdateMission(image *multipart.FileHeader, missionID string, data Mission) error
	UpdateMissionStage(MissionStageID string, data MissionStage) error
}
