package request

import "recything/features/mission/entity"

func MissionRequestToMissionCore(data Mission) entity.Mission {
	missionCore := entity.Mission{
		Title:       data.Title,
		Point:       data.Point,
		Description: data.Description,
		StartDate:   data.Start_Date,
		EndDate:     data.End_Date,
	}

	return missionCore
}

func StagesRequestToStagesCore(missionID string, stage Stage) entity.MissionStage {
	missionStagesCore := entity.MissionStage{
		MissionID:   missionID,
		Title:       stage.Title,
		Description: stage.Description,
	}
	return missionStagesCore
}

func ListMissiStagesRequestToMissiStagesCore(data MissionStages) []entity.MissionStage {
	missionStagesCore := []entity.MissionStage{}
	for _, stage := range data.Stages {
		result := StagesRequestToStagesCore(data.MissionID, stage)
		missionStagesCore = append(missionStagesCore, result)
	}
	return missionStagesCore
}