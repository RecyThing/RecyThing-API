package request

import "recything/features/trash_exchange/entity"

func TrashExchangeDetailsRequestToTrashExchangeDetailsCore(data TrashExchangeDetailRequest) entity.TrashExchangeDetailCore {
	return entity.TrashExchangeDetailCore{
		TrashType:   data.TrashType,
		Unit:        data.Unit,
	}
}

func ListTrashExchangeDetailsRequestToTrashExchangeDetailsCore(data []TrashExchangeDetailRequest) []entity.TrashExchangeDetailCore {
	listTrashExchange := []entity.TrashExchangeDetailCore{}
	for _, v := range data {
		trashExchange := TrashExchangeDetailsRequestToTrashExchangeDetailsCore(v)
		listTrashExchange = append(listTrashExchange, trashExchange)
	}

	return listTrashExchange
}

func TrashExchangeRequestToTrashExchangeCore(data TrashExchangeRequest) entity.TrashExchangeCore {
	TrashExchangeCore := entity.TrashExchangeCore{
		Name:      data.Name,
		EmailUser: data.EmailUser,
		Address:   data.Address,
	}
	trashExchange := ListTrashExchangeDetailsRequestToTrashExchangeDetailsCore(data.TrashExchangeDetails)
	TrashExchangeCore.TrashExchangeDetails = trashExchange
	return TrashExchangeCore
}