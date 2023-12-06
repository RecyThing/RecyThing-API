package entity

import (
	"mime/multipart"
	"recything/utils/helper"
	"recything/utils/pagination"
)

type MissionRepositoryInterface interface {
	CreateMission(input Mission) error
	FindAllMission(page, limit int, search, filter string) ([]Mission, pagination.PageInfo, int, error)
	FindById(missionID string) (Mission, error)
	GetCountMission(status, search string) (int, error)
	GetCountMissionApproval(filter, search string) (int, error)
	GetCountDataMissionApproval() (helper.CountMissionApproval, error)
	GetImageURL(missionID string) (string, error)
	UpdateMission(missionID string, data Mission) error
	DeleteMission(missionID string) error

	UpdateMissionStage(missionStageID string, data []MissionStage) error
	ClaimMission(userID string, data ClaimedMission) error
	FindClaimed(userID, missionID string) error

	FindUploadMissionStatus(id, missionID, userID, status string) error
	FindUploadById(id string) error
	CreateUploadMissionTask(userID string, data UploadMissionTaskCore, images []*multipart.FileHeader) error
	UpdateUploadMissionTask(id string, images []*multipart.FileHeader, data UploadMissionTaskCore) error

	FindAllMissionApproval(page, limit int, search, filter string) ([]UploadMissionTaskCore, pagination.PageInfo, int, error)
	FindMissionApprovalById(UploadMissionTaskID string) (UploadMissionTaskCore, error)
	UpdateStatusMissionApproval(uploadMissionTaskID, status, reason string) error
}

type MissionServiceInterface interface {
	CreateMission(image *multipart.FileHeader, data Mission) error
	// CreateMission(data Mission) error
	FindAllMission(page, limit, search, filter string) ([]Mission, pagination.PageInfo, int, error)
	FindById(missionID string) (Mission, error)
	UpdateMission(image *multipart.FileHeader, missionID string, data Mission) error

	UpdateMissionStage(missionStageID string, data []MissionStage) error
	ClaimMission(userID string, data ClaimedMission) error
	DeleteMission(missionID string) error

	CreateUploadMissionTask(userID string, data UploadMissionTaskCore, images []*multipart.FileHeader) error
	UpdateUploadMissionTask(userID, id string, images []*multipart.FileHeader, data UploadMissionTaskCore) error

	FindAllMissionApproval(page, limit, search, filter string) ([]UploadMissionTaskCore, pagination.PageInfo, int, error)
	FindMissionApprovalById(UploadMissionTaskID string) (UploadMissionTaskCore, error)
	UpdateStatusMissionApproval(uploadMissionTaskID, status, reason string) error
}
