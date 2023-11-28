package request

import (
	"recything/features/mission/entity"
)

type Mission struct {
	Title string `form:"title"`
	MissionImage  string          `form:"mission_image"`
	Point         int             `form:"point"`
	Description   string          `form:"description"`
	StartDate     string          `form:"start_date"`
	EndDate       string          `form:"end_date"`
	MissionStages []MissionStage `form:"mission_stages"`
}

type MissionStage struct {
	Title       string `form:"title"`
	Description string `form:"description"`
}

func MissionRequestToMissionCore(missi Mission) entity.Mission {
	missionCore := entity.Mission{
		Title: missi.Title,
		// Creator: missi.Creator,
		MissionImage: missi.MissionImage,
		Point:        missi.Point,
		Description:  missi.Description,
		StartDate:    missi.StartDate,
		EndDate:      missi.EndDate,
	}
	missionStagesCore := ListMissionStagesRequestToMissionStagesCore(missi.MissionStages)
	missionCore.MissionStages = missionStagesCore
	return missionCore
}

func MissionStagesRequestToMissionStagesCore(missionStages MissionStage) entity.MissionStages {
	missionStagesCore := entity.MissionStages{
		Title:       missionStages.Title,
		Description: missionStages.Description,
	}
	return missionStagesCore
}

func ListMissionStagesRequestToMissionStagesCore(missionStages []MissionStage) []entity.MissionStages {
	missionStagesCore := []entity.MissionStages{}
	for _, misiStages := range missionStages {
		missi := MissionStagesRequestToMissionStagesCore(misiStages)
		missionStagesCore = append(missionStagesCore, missi)
	}
	return missionStagesCore
}
