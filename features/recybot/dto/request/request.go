package request

import "recything/features/recybot/entity"

type RecybotRequest struct {
	Category string `json:"category" form:"category"`
	Question string `json:"question" form:"question"`
}

func RequestRecybotToCoreRecybot(recybot RecybotRequest) entity.RecybotCore {
	return entity.RecybotCore{
		Category: recybot.Category,
		Question: recybot.Category,
	}
}
