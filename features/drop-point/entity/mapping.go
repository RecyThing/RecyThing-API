package entity

import "recything/features/drop-point/model"

func OperationalSchedulesModelToOperationalSchedulesCore(data model.OperationalSchedules) OperationalSchedulesCore {
	return OperationalSchedulesCore{
		Id:          data.Id,
		DropPointId: data.DropPointId,
		Days:        data.Days,
		Open:        data.Open,
		Close:       data.Close,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func ListOperationalSchedulesModelToOperationalSchedulesCore(data []model.OperationalSchedules) []OperationalSchedulesCore {
	coreOperational := []OperationalSchedulesCore{}
	for _, v := range data {
		operational := OperationalSchedulesModelToOperationalSchedulesCore(v)
		coreOperational = append(coreOperational, operational)
	}
	return coreOperational
}

func DropPointModelToDropPointCore(data model.DropPoint) DropPointCore {
	dropPointCore := DropPointCore{
		Id:        data.Id,
		Name:      data.Name,
		Address:   data.Address,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	operational := ListOperationalSchedulesModelToOperationalSchedulesCore(data.OperationalSchedules)
	dropPointCore.OperationalSchedules = operational
	return dropPointCore
}

func OperationalSchedulesCoreToOperationalSchedulesModel(data OperationalSchedulesCore) model.OperationalSchedules {
	return model.OperationalSchedules{
		Id:          data.Id,
		DropPointId: data.DropPointId,
		Days:        data.Days,
		Open:        data.Open,
		Close:       data.Close,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func ListOperationalSchedulesCoreToOperationalSchedulesModel(data []OperationalSchedulesCore) []model.OperationalSchedules {
	coreOperational := []model.OperationalSchedules{}
	for _, v := range data {
		operational := OperationalSchedulesCoreToOperationalSchedulesModel(v)
		coreOperational = append(coreOperational, operational)
	}
	return coreOperational
}

func DropPointCoreToDropPointModel(data DropPointCore) model.DropPoint {
	dropPointModel := model.DropPoint{
		Id:        data.Id,
		Name:      data.Name,
		Address:   data.Address,
		Latitude:  data.Latitude,
		Longitude: data.Longitude,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	dropPoint := ListOperationalSchedulesCoreToOperationalSchedulesModel(data.OperationalSchedules)
	dropPointModel.OperationalSchedules = dropPoint
	return dropPointModel
}
