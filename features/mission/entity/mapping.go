package entity

import (
	"recything/features/mission/model"
	"time"
)

func MissionCoreToMissionModel(data Mission) model.Mission {
	missionModel := model.Mission{
		Title:        data.Title,
		Status:       data.Status,
		AdminID:      data.AdminID,
		MissionImage: data.MissionImage,
		Point:        data.Point,
		Description:  data.Description,
		StartDate:    data.StartDate,
		EndDate:      data.EndDate,
	}
	missionStagesModel := ListMissionStagesCoreToMissionStagesModel(data.MissionStages)
	missionModel.MissionStages = missionStagesModel
	return missionModel
}

func MissionHistoriesCoreToMap(data MissionHistories) map[string]interface{} {
	return map[string]interface{}{
		"id_transaction":   data.TransactionID,
		"point":            data.Point,
		"type_transaction": "hadiah mission",
		"time_transaction": data.CreatedAt.Format("15:04:05.000"),
		"created_at":       data.CreatedAt.Format(time.RFC3339),
	}
}

func MissionHistoriesCoreToMapDetail(data MissionHistories) map[string]interface{} {
	return map[string]interface{}{
		"id_transaction":   data.TransactionID,
		"mission_id":       data.MissionID,
		"title":            data.Title,
		"status":           data.StatusApproval,
		"point":            data.Point,
		"type_transaction": "reward hadiah mission",
		"time_transaction": data.CreatedAt.Format("15:04:05.000"),
		"created_at":       data.CreatedAt.Format(time.RFC3339),
	}
}

func UploadTaskModelToMissionHistoriesCore(data model.UploadMissionTask, dataMission model.Mission) MissionHistories {
	return MissionHistories{
		TransactionID:  data.ID,
		MissionID:      data.MissionID,
		Title:          dataMission.Title,
		StatusApproval: data.Status,
		Point:          dataMission.Point,
		CreatedAt:      data.CreatedAt,
	}
}

func MissionModelToMissionCore(data model.Mission) Mission {
	missionCore := Mission{
		ID:           data.ID,
		Title:        data.Title,
		Status:       data.Status,
		AdminID:      data.AdminID,
		MissionImage: data.MissionImage,
		Point:        data.Point,
		Description:  data.Description,
		StartDate:    data.StartDate,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
		EndDate:      data.EndDate,
	}
	missionStagesCore := listMissionStagesModelToMissionStagesCore(data.MissionStages)
	missionCore.MissionStages = missionStagesCore
	return missionCore
}

func MissionStagesCoreToMissionStagesModel(data MissionStage) model.MissionStage {
	missionStagesModel := model.MissionStage{
		MissionID:   data.MissionID,
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
	}
	return missionStagesModel
}

func MissionStagesModelToMissionStagesCore(data model.MissionStage) MissionStage {
	missionStagesCore := MissionStage{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
	}
	return missionStagesCore
}

func listMissionStagesModelToMissionStagesCore(data []model.MissionStage) []MissionStage {
	missionStagesCore := []MissionStage{}
	for _, misiStages := range data {
		result := MissionStagesModelToMissionStagesCore(misiStages)
		missionStagesCore = append(missionStagesCore, result)
	}
	return missionStagesCore
}
func ListMissionStagesCoreToMissionStagesModel(data []MissionStage) []model.MissionStage {
	missionStagesModel := []model.MissionStage{}
	for _, misiStages := range data {
		result := MissionStagesCoreToMissionStagesModel(misiStages)
		missionStagesModel = append(missionStagesModel, result)
	}
	return missionStagesModel
}

func ListMissionModelToMissionCore(data []model.Mission) []Mission {
	missions := []Mission{}
	for _, mission := range data {
		result := MissionModelToMissionCore(mission)
		missions = append(missions, result)
	}
	return missions
}

func ListMissionCoreToMissionMission(data []Mission) []model.Mission {
	missions := []model.Mission{}
	for _, mission := range data {
		result := MissionCoreToMissionModel(mission)
		missions = append(missions, result)
	}
	return missions
}

// claimed mission

func ClaimedCoreToClaimedMissionModel(data ClaimedMission) model.ClaimedMission {
	return model.ClaimedMission{
		UserID:    data.UserID,
		MissionID: data.MissionID,
		Claimed:   data.Claimed,
	}

}

func UploadMissionTaskCoreToUploadMissionTaskModel(data UploadMissionTaskCore) model.UploadMissionTask {
	return model.UploadMissionTask{
		UserID:      data.UserID,
		MissionID:   data.MissionID,
		Description: data.Description,
		Images:      ListImageUploadMissionCoreToImageUploadMissionModel(data.Images),
		Status:      data.Status,
	}
}

func UploadMissionTaskModelToUploadMissionTaskCore(data model.UploadMissionTask) UploadMissionTaskCore {
	return UploadMissionTaskCore{
		ID:          data.ID,
		UserID:      data.UserID,
		MissionID:   data.MissionID,
		Description: data.Description,
		Reason:      data.Reason,
		Images:      ListImageUploadMissionModelToImageUploadMissionCore(data.Images),
		Status:      data.Status,
		CreatedAt:   data.CreatedAt,
	}
}

func ImageUploadMissionCoreToImageUploadMissionModel(data ImageUploadMissionCore) model.ImageUploadMission {
	return model.ImageUploadMission{
		UploadMissionTaskID: data.UploadMissionTaskID,
		Image:               data.Image,
	}
}

func ImageUploadMissionModelToImageUploadMissionCore(data model.ImageUploadMission) ImageUploadMissionCore {
	return ImageUploadMissionCore{
		ID:                  data.ID,
		UploadMissionTaskID: data.UploadMissionTaskID,
		Image:               data.Image,
		CreatedAt:           data.CreatedAt,
	}
}

func ListUploadMissionTaskModelToUploadMissionTaskCore(data []model.UploadMissionTask) []UploadMissionTaskCore {
	dataTask := []UploadMissionTaskCore{}
	for _, v := range data {
		result := UploadMissionTaskModelToUploadMissionTaskCore(v)
		dataTask = append(dataTask, result)
	}
	return dataTask

}

func ListImageUploadMissionModelToImageUploadMissionCore(data []model.ImageUploadMission) []ImageUploadMissionCore {
	dataImage := []ImageUploadMissionCore{}
	for _, v := range data {
		result := ImageUploadMissionModelToImageUploadMissionCore(v)
		dataImage = append(dataImage, result)
	}
	return dataImage
}

func ListImageUploadMissionCoreToImageUploadMissionModel(data []ImageUploadMissionCore) []model.ImageUploadMission {
	dataImage := []model.ImageUploadMission{}
	for _, v := range data {
		result := ImageUploadMissionCoreToImageUploadMissionModel(v)
		dataImage = append(dataImage, result)
	}
	return dataImage
}

func MissionModelTomissionHistoriesCore(data model.Mission) MissionHistories {
	return MissionHistories{
		MissionID:     data.ID,
		Title:         data.Title,
		StatusMission: data.Status,
		MissionImage:  data.MissionImage,
		Point:         data.Point,
		Description:   data.Description,
		StartDate:     data.StartDate,
		EndDate:       data.EndDate,
		MissionStages: []MissionStage{},
		CreatedAt:     data.CreatedAt,
	}
}

func ListMissionModelTomissionHistoriesCore(data []model.Mission) []MissionHistories {
	dataMissi := []MissionHistories{}
	for _, v := range data {
		result := MissionModelTomissionHistoriesCore(v)
		dataMissi = append(dataMissi, result)
	}
	return dataMissi
}

func MissionToMissionHistoriesCore(data model.Mission, claimed model.ClaimedMission, upMisTask model.UploadMissionTask) MissionHistories {
	return MissionHistories{
		MissionID:      data.ID,
		ClaimedID:      claimed.ID,
		TransactionID:  upMisTask.ID,
		Title:          data.Title,
		StatusApproval: upMisTask.Status,
		StatusMission:  data.Status,
		MissionImage:   data.MissionImage,
		Reason:         upMisTask.Reason,
		Point:          data.Point,
		Description:    data.Description,
		StartDate:      data.StartDate,
		EndDate:        data.EndDate,
		CreatedAt:      upMisTask.CreatedAt,
	}
}
