package request

import (
	"recything/features/drop-point/entity"
)

func OperationalSchedulesRequestToOperationalSchedulesCore(data OperationalSchedulesRequest) entity.OperationalSchedulesCore {
	return entity.OperationalSchedulesCore{
		Days:  data.Days,
		Open:  data.Open,
		Close: data.Close,
	}
}

func ListOperationalSchedulesRequestToOperationalSchedulesCore(data []OperationalSchedulesRequest) []entity.OperationalSchedulesCore {
	listDropPoint := []entity.OperationalSchedulesCore{}
	for _, v := range data {
		dropPoint := OperationalSchedulesRequestToOperationalSchedulesCore(v)
		listDropPoint = append(listDropPoint, dropPoint)
	}

	return listDropPoint
}

func DropPointRequestToReportCore(data DropPointRequest) entity.DropPointCore {
	reportCore := entity.DropPointCore{
		Name:      data.Name,
		Address:   data.Address,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
	}
	dropPoint := ListOperationalSchedulesRequestToOperationalSchedulesCore(data.OperationalSchedules)
	reportCore.OperationalSchedules = dropPoint
	return reportCore
}
