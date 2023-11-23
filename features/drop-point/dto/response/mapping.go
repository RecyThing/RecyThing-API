package response

import "recything/features/drop-point/entity"

func OperationalSchedulesCoreToOperationalSchedulesResponse(data entity.OperationalSchedulesCore) OperationalSchedulesResponse {
	return OperationalSchedulesResponse{
		Days:  data.Days,
		Open:  data.Open,
		Close: data.Close,
	}
}

func ListOperationalSchedulesCoreToOperationalSchedulesResponse(data []entity.OperationalSchedulesCore) []OperationalSchedulesResponse {
	responseOperational := []OperationalSchedulesResponse{}
	for _, v := range data {
		operational := OperationalSchedulesCoreToOperationalSchedulesResponse(v)
		responseOperational = append(responseOperational, operational)
	}
	return responseOperational
}

func DropPointCoreToDropPointResponse(data entity.DropPointCore) DropPointResponse {
	dropPointResponse := DropPointResponse{
		Name:      data.Name,
		Address:   data.Address,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
	}

	dropPoint := ListOperationalSchedulesCoreToOperationalSchedulesResponse(data.OperationalSchedules)
	dropPointResponse.OperationalSchedules = dropPoint
	return dropPointResponse
}
