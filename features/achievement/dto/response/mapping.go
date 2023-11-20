package response

import "recything/features/achievement/entity"

func AchievementCoreToAchievementResponse(data entity.AchievementCore) AchievementResponse {
	return AchievementResponse{
		Name:        data.Name,
		TargetPoint: data.TargetPoint,
		TotalUser:   data.TotalUser,
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
