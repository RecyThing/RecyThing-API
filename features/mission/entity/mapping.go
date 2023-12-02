package entity

import (
	"recything/features/mission/model"
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
		Title:       data.Title,
		Description: data.Description,
	}
	return missionStagesModel
}

func MissionStagesModelToMissionStagesCore(data model.MissionStage) MissionStage {
	missionStagesCore := MissionStage{
		Title:       data.Title,
		Description: data.Description,
	}
	return missionStagesCore
}

// func StageCoreToMissionStageModel(data Stage) model.MissionStage {
// 	missionStagesModel := model.MissionStage{
// 		Title:       data.Title,
// 		Description: data.Description,
// 	}
// 	return missionStagesModel
// }

func listMissionStagesModelToMissionStagesCore(data []model.MissionStage) []MissionStage {
	missionStagesCore := []MissionStage{}
	for _, misiStages := range data {
		result := MissionStagesModelToMissionStagesCore(misiStages)
		missionStagesCore = append(missionStagesCore, result)
	}
	return missionStagesCore
}

func ListMissionModelToMissionCore(data []model.Mission) []Mission {
	missions := []Mission{}
	for _, mission := range data {
		result := MissionModelToMissionCore(mission)
		missions = append(missions, result)
	}
	return missions
}
func ListMissionStagesCoreToMissionStagesModel(data []MissionStage) []model.MissionStage {
	missionStagesModel := []model.MissionStage{}
	for _, misiStages := range data {
		result := MissionStagesCoreToMissionStagesModel(misiStages)
		missionStagesModel = append(missionStagesModel, result)
	}
	return missionStagesModel
}
