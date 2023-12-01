package entity

import (
	"mime/multipart"
	"recything/utils/pagination"
)

type MissionRepositoryInterface interface {
	CreateMission(input Mission) error
	FindAllMission(page, limit int, search, filter string) ([]Mission, pagination.PageInfo, int, error)
	GetCount(filter, search string) (int, error)
	CreateMissionStages(input []MissionStage) error
	GetAdminIDbyMissionID(missionID string) (string, error)
	SaveChangesStatusMission(data Mission) error
	UpdateMission(missionID string, data Mission) error 
	UpdateMissionStage(missionStageID string, data Stage) error
	GetById(missionID string) (Mission, error)
}

type MissionServiceInterface interface {
	CreateMission(image *multipart.FileHeader, data Mission) error
	FindAllMission(page, limit, search, filter string) ([]Mission, pagination.PageInfo, int, error)
	ChangesStatusMission(data Mission) error
	UpdateMission(image *multipart.FileHeader, missionID string, data Mission) error
	CreateMissionStages(adminID, missionID string, data []MissionStage) error
	UpdateMissionStage(missionStageID string, data Stage) error

}
