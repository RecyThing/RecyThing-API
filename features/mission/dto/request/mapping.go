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

func AddMissionStageToMissionStageCore(addMissionStage AddMissionStage) []entity.MissionStage {
    var missionStages []entity.MissionStage
    for _, stage := range addMissionStage.Stages {
        newStage := entity.MissionStage{
            MissionID:   addMissionStage.MissionID,
            Title:       stage.Title,
            Description: stage.Description,
           
        }
        missionStages = append(missionStages, newStage)
    }
    return missionStages
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


func ClaimRequestToClaimCore( claim Claim) entity.ClaimedMission {
	return entity.ClaimedMission {
		MissionID: claim.MissionID,
	}
	
}