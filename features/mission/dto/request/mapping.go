package request

import "recything/features/mission/entity"

func MissionRequestToMissionCore(missi Mission) entity.Mission {
	missionCore := entity.Mission{
		Title:        missi.Title,
		MissionImage: missi.MissionImage,
		Point:        missi.Point,
		Description:  missi.Description,
		StartDate:    missi.Start_Date,
		EndDate:      missi.End_Date,
	}

	missionStagesCore := ListMissionStagesRequestToMissionStagesCore(missi.MissionStages)
	missionCore.MissionStages = missionStagesCore
	return missionCore
}

func MissionStagesRequestToMissionStagesCore(missionStages MissionStage) entity.MissionStage {
	missionStagesCore := entity.MissionStage{
		Title:       missionStages.Title,
		Description: missionStages.Description,
	}

	return missionStagesCore
}

func ListMissionStagesRequestToMissionStagesCore(missionStages []MissionStage) []entity.MissionStage {
	missionStagesCore := []entity.MissionStage{}
	for _, misiStages := range missionStages {
		missi := MissionStagesRequestToMissionStagesCore(misiStages)
		missionStagesCore = append(missionStagesCore, missi)
	}
	return missionStagesCore
}
