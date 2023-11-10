package request

import "recything/features/recybot/entity"

func RequestRecybotToCoreRecybot(recybot RecybotRequest) entity.RecybotCore {
	return entity.RecybotCore{
		Category: recybot.Category,
		Question: recybot.Category,
	}
}

func ListRequestRecybotToCoreRecybot(recybot []RecybotRequest) []entity.RecybotCore {
	list := []entity.RecybotCore{}
	for _, v := range recybot {
		result := RequestRecybotToCoreRecybot(v)
		list = append(list, result)
	}
	return list
}
