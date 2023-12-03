package entity

import (
	"mime/multipart"
	"recything/utils/pagination"
)

type MissionRepositoryInterface interface {
	CreateMission(input Mission) error
	FindAllMission(page, limit int, search, status string) ([]Mission, pagination.PageInfo, int, error)
	FindById(missionID string) (Mission, error)
	GetCount(status, search string) (int, error)
	SaveChangesStatusMission(data Mission) error
	UpdateMission(missionID string, data Mission) error
	DeleteMission(missionID string) error


	UpdateMissionStage(missionStageID string, data MissionStage) error
	GetImageURL(missionID string)(string, error)
	ClaimMission(userID string, data ClaimedMission) error
	FindClaimed(userID, missionID string) error
}

type MissionServiceInterface interface {
	CreateMission(image *multipart.FileHeader, data Mission) error
	FindAllMission(page, limit, search, status string) ([]Mission, pagination.PageInfo, int, error)
	FindById(missionID string) (Mission, error)
	UpdateMission(image *multipart.FileHeader, missionID string, data Mission) error

	UpdateMissionStage(MissionStageID string, data MissionStage) error
	ClaimMission(userID string, data ClaimedMission) error
	DeleteMission(missionID string) error
}
