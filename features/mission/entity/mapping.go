package entity

import (
	"recything/features/mission/model"
)

func MissionCoreToMissionModel(missi Mission) model.Mission {
	missionModel := model.Mission{
		Title:        missi.Title,
		Status:       missi.Status,
		AdminID:      missi.AdminID,
		MissionImage: missi.MissionImage,
		Point:        missi.Point,
		Description:  missi.Description,
		StartDate:    missi.StartDate,
		EndDate:      missi.EndDate,
	}
	missionStagesModel := ListMissionStagesCoreToMissionStagesModel(missi.MissionStages)
	missionModel.MissionStages = missionStagesModel
	return missionModel
}

func MissionStagesCoreToMissionStagesModel(missionStages MissionStage) model.MissionStage {
	missionStagesModel := model.MissionStage{
		Title:       missionStages.Title,
		Description: missionStages.Description,
	}
	return missionStagesModel
}

func ListMissionStagesCoreToMissionStagesModel(missionStages []MissionStage) []model.MissionStage {
	missionStagesModel := []model.MissionStage{}
	for _, misiStages := range missionStages {
		missi := MissionStagesCoreToMissionStagesModel(misiStages)
		missionStagesModel = append(missionStagesModel, missi)
	}
	return missionStagesModel
}

func MissionModelToMissionCore(missi model.Mission) Mission {
	missionCore := Mission{
		ID:           missi.ID,
		Title:        missi.Title,
		Status:       missi.Status,
		AdminID:      missi.AdminID,
		MissionImage: missi.MissionImage,
		Point:        missi.Point,
		Description:  missi.Description,
		StartDate:    missi.StartDate,
		EndDate:      missi.EndDate,
	}
	missionStagesCore := listMissionStagesModelToMissionStagesCore(missi.MissionStages)
	missionCore.MissionStages = missionStagesCore
	return missionCore
}

func MissionStagesModelToMissionStagesCore(missionStages model.MissionStage) MissionStage {
	missionStagesCore := MissionStage{
		Title:       missionStages.Title,
		Description: missionStages.Description,
	}
	return missionStagesCore
}

func listMissionStagesModelToMissionStagesCore(mission []model.MissionStage) []MissionStage {
	missionStagesCore := []MissionStage{}
	for _, misiStages := range mission {
		missi := MissionStagesModelToMissionStagesCore(misiStages)
		missionStagesCore = append(missionStagesCore, missi)
	}
	return missionStagesCore
}

func ListMissionModelToMissionCore(mission []model.Mission) []Mission {
	missions := []Mission{}
	for _, mission := range mission {
		missionCore := MissionModelToMissionCore(mission)
		missions = append(missions, missionCore)
	}
	return missions
}
