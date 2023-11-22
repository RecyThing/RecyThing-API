package response

import "recything/features/achievement/entity"

func AchievementCoreToAchievementResponse(data entity.AchievementCore) AchievementResponse {
	return AchievementResponse{
		Id:           data.Id,
		Name:         data.Name,
		TargetPoint:  data.TargetPoint,
		TotalClaimed: data.TotalClaimed,
	}
}

func ListAchievementCoreToAchievementResponse(data []entity.AchievementCore) []AchievementResponse {
	dataAchievement := []AchievementResponse{}
	for _, achievement := range data {
		achievementRespon := AchievementCoreToAchievementResponse(achievement)
		dataAchievement = append(dataAchievement, achievementRespon)
	}
	return dataAchievement
}
