package service

import (
	"errors"
	"log"
	"recything/features/drop-point/entity"
	"recything/utils/constanta"
	"recything/utils/pagination"
	"recything/utils/validation"
)

type dropPointService struct {
	dropPointRepository entity.DropPointRepositoryInterface
}

func NewDropPointService(dropPoint entity.DropPointRepositoryInterface) entity.DropPointServiceInterface {
	return &dropPointService{
		dropPointRepository: dropPoint,
	}
}

// Create implements entity.DropPointServiceInterface.
func (dps *dropPointService) CreateDropPoint(data entity.DropPointCore) (entity.DropPointCore, error) {
	if len(data.OperationalSchedules) > 7 {
		return entity.DropPointCore{}, errors.New("jumlah hari tidak boleh lebih dari 7")
	}

	uniqueDays := make(map[string]bool)
	for _, schedule := range data.OperationalSchedules {
		switch schedule.Days {
		case "senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu":
			// Validasi apakah hari sudah pernah digunakan
			if uniqueDays[schedule.Days] {
				return entity.DropPointCore{}, errors.New("hari tidak boleh sama")
			}
			uniqueDays[schedule.Days] = true

			// Validasi waktu buka dan tutup
			if err := validation.ValidateTime(schedule.Open, schedule.Close); err != nil {
				return entity.DropPointCore{}, err
			}
		default:
			return entity.DropPointCore{}, errors.New("hari tidak valid")
		}
	}

	for _, day := range []string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"} {
		if !uniqueDays[day] {
			data.OperationalSchedules = append(data.OperationalSchedules, entity.OperationalSchedulesCore{
				Days:  day,
				Open:  "tutup",
				Close: "tutup",
			})
		}
	}

	dataDropPoint, err := dps.dropPointRepository.CreateDropPoint(data)
	if err != nil {
		return entity.DropPointCore{}, err
	}

	return dataDropPoint, nil
}

// DeleteDropPoint implements entity.DropPointServiceInterface.
func (dps *dropPointService) DeleteDropPointById(id string) error {
	if id == "" {
		return errors.New(constanta.ERROR_ID_INVALID)
	}

	err := dps.dropPointRepository.DeleteDropPointById(id)
	if err != nil {
		return err
	}

	return nil
}

// GetAllDropPoint implements entity.DropPointServiceInterface.
func (dps *dropPointService) GetAllDropPoint(page, limit int, name, address string) ([]entity.DropPointCore, pagination.PageInfo, error) {
	if limit > 10 {
        return nil, pagination.PageInfo{}, errors.New("limit tidak boleh lebih dari 10")
    }

	page, limit = validation.ValidatePaginationParameters(page, limit)
	
	dropPointCores, pageInfo, err := dps.dropPointRepository.GetAllDropPoint(page, limit, name, address)
	if err != nil {
		return nil, pagination.PageInfo{}, err
	}

	return dropPointCores, pageInfo, nil
}

// GetById implements entity.DropPointServiceInterface.
func (dps *dropPointService) GetDropPointById(id string) (entity.DropPointCore, error) {
	if id == "" {
		return entity.DropPointCore{}, errors.New(constanta.ERROR_ID_INVALID)
	}

	idDropPoint, err := dps.dropPointRepository.GetDropPointById(id)
	if err != nil {
		return entity.DropPointCore{}, err
	}
	
	return idDropPoint, err
}

// UpdateById implements entity.DropPointServiceInterface.
func (dps *dropPointService) UpdateDropPointById(id string, data entity.DropPointCore) (entity.DropPointCore, error) {
	if id == "" {
		return entity.DropPointCore{}, errors.New(constanta.ERROR_ID_INVALID)
	}

	if len(data.OperationalSchedules) > 7 {
		return entity.DropPointCore{}, errors.New("jumlah hari tidak boleh lebih dari 7")
	}

	uniqueDays := make(map[string]bool)
	for _, schedule := range data.OperationalSchedules {
		switch schedule.Days {
		case "senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu":
			// Validasi apakah hari sudah pernah digunakan
			if uniqueDays[schedule.Days] {
				log.Println("Hari sudah digunakan:", schedule.Days)
				return entity.DropPointCore{}, errors.New("hari tidak boleh sama")
			}
			uniqueDays[schedule.Days] = true

			// Validasi waktu buka dan tutup
			if err := validation.ValidateTime(schedule.Open, schedule.Close); err != nil {
				log.Println("Validasi waktu buka dan tutup gagal:", err)
				return entity.DropPointCore{}, err
			}
		default:
			return entity.DropPointCore{}, errors.New("hari tidak valid")
		}
	}

	updatedData, err := dps.dropPointRepository.UpdateDropPointById(id, data)
	if err != nil {
		return entity.DropPointCore{}, err
	}

	return updatedData, nil
}
