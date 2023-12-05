package service

import (
	"errors"
	dropPoint "recything/features/drop-point/entity"
	trashCategory "recything/features/trash_category/entity"
	trashExchange "recything/features/trash_exchange/entity"
	user "recything/features/user/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/pagination"
	"recything/utils/validation"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type trashExchangeService struct {
	trashExchangeRepository trashExchange.TrashExchangeRepositoryInterface
	dropPointRepository     dropPoint.DropPointRepositoryInterface
	userRepository          user.UsersRepositoryInterface
	trashCategoryRepository trashCategory.TrashCategoryRepositoryInterface
}

func NewTrashExchangeService(trashExchange trashExchange.TrashExchangeRepositoryInterface, dropPoint dropPoint.DropPointRepositoryInterface, user user.UsersRepositoryInterface, trashCategory trashCategory.TrashCategoryRepositoryInterface) trashExchange.TrashExchangeServiceInterface {
	return &trashExchangeService{
		trashExchangeRepository: trashExchange,
		dropPointRepository:     dropPoint,
		userRepository:          user,
		trashCategoryRepository: trashCategory,
	}
}

// CreateTrashExchange implements entity.TrashExchangeServiceInterface.
func (tes *trashExchangeService) CreateTrashExchange(data trashExchange.TrashExchangeCore) (trashExchange.TrashExchangeCore, error) {
	
	data.Id = helper.GenerateRandomID(4)
	errEmpty := validation.CheckDataEmpty(data.Name, data.EmailUser, data.Address)
	if errEmpty != nil {
		return trashExchange.TrashExchangeCore{}, errEmpty
	}

	user, err := tes.userRepository.FindByEmail(data.EmailUser)
	if err != nil {
		return trashExchange.TrashExchangeCore{}, errors.New("pengguna dengan email tersebut tidak ditemukan")
	}

	dropPoint, err := tes.dropPointRepository.GetDropPointByAddress(data.Address)
	if err != nil {
		return trashExchange.TrashExchangeCore{}, errors.New("alamat drop point tidak ditemukan")
	}
	data.DropPointId = dropPoint.Id

	totalPoints := 0
	totalUnits := 0.0
	var details []trashExchange.TrashExchangeDetailCore
	for _, detail := range data.TrashExchangeDetails {

		errEmptyDetail := validation.CheckDataEmpty(detail.TrashType, detail.Unit)
		if errEmptyDetail != nil {
			return trashExchange.TrashExchangeCore{}, errEmptyDetail
		}

		titleCase := cases.Title(language.Indonesian)
		detail.TrashType = titleCase.String(detail.TrashType)
		trashCategory, err := tes.trashCategoryRepository.GetByType(detail.TrashType)
		if err != nil {
			return trashExchange.TrashExchangeCore{}, errors.New("kategori sampah tidak ditemukan")
		}

		detail.Unit = titleCase.String(detail.Unit)
		unit, err := helper.ConvertUnitToDecimal(detail.Unit)
		if err != nil {
			return trashExchange.TrashExchangeCore{}, err
		}

		detail.TotalPoints = int(unit * float64(trashCategory.Point))
		totalPoints += detail.TotalPoints

		totalUnits += unit
		details = append(details, detail)
	}

	data.TotalPoint = totalPoints
	data.TotalUnit = totalUnits

	user.Point += data.TotalPoint
	// Update user
	err = tes.userRepository.UpdateById(user.Id, user)
	if err != nil {
		return trashExchange.TrashExchangeCore{}, errors.New("gagal memperbarui nilai point pengguna")
	}

	result, err := tes.trashExchangeRepository.CreateTrashExchange(data)
	if err != nil {
		return trashExchange.TrashExchangeCore{}, errors.New("gagal menyimpan data trash exchange")
	}

	for _, detail := range details {
		detail.TrashExchangeId = result.Id
		_, err := tes.trashExchangeRepository.CreateTrashExchangeDetails(detail)
		if err != nil {
			return trashExchange.TrashExchangeCore{}, errors.New("gagal menyimpan data trash exchange detail")
		}
	}

	return result, nil
}

// DeleteTrashExchangeById implements entity.TrashExchangeServiceInterface.
func (tes *trashExchangeService) DeleteTrashExchangeById(id string) error {
	if id == "" {
		return errors.New(constanta.ERROR_ID_INVALID)
	}

	err := tes.trashExchangeRepository.DeleteTrashExchangeById(id)
	if err != nil {
		return err
	}

	return nil
}

// GetAllTrashExchange implements entity.TrashExchangeServiceInterface.
func (tes *trashExchangeService) GetAllTrashExchange(page, limit, search string) ([]trashExchange.TrashExchangeCore, pagination.PageInfo, int, error) {
	pageInt, limitInt, err := validation.ValidateTypePaginationParameter(limit, page)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	pageValid, limitValid := validation.ValidateCountLimitAndPage(pageInt, limitInt)

	dropPointCores, pageInfo, count, err := tes.trashExchangeRepository.GetAllTrashExchange(pageValid, limitValid, search)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	return dropPointCores, pageInfo, count, nil
}

// GetTrashExchangeById implements entity.TrashExchangeServiceInterface.
func (tes *trashExchangeService) GetTrashExchangeById(id string) (trashExchange.TrashExchangeCore, error) {
	if id == "" {
		return trashExchange.TrashExchangeCore{}, errors.New(constanta.ERROR_ID_INVALID)
	}

	idtrashExchange, err := tes.trashExchangeRepository.GetTrashExchangeById(id)
	if err != nil {
		return trashExchange.TrashExchangeCore{}, err
	}

	return idtrashExchange, err
}
