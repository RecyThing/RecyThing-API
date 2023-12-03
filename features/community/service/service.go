package service

import (
	"errors"
	"mime/multipart"
	"recything/features/community/entity"
	"recything/utils/constanta"
	"recything/utils/pagination"
	"recything/utils/validation"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type communityService struct {
	communityRepository entity.CommunityRepositoryInterface
}

func NewCommunityService(community entity.CommunityRepositoryInterface) entity.CommunityServiceInterface {
	return &communityService{
		communityRepository: community,
	}
}

// CreateCommunity implements entity.CommunityServiceInterface.
func (cs *communityService) CreateCommunity(image *multipart.FileHeader, data entity.CommunityCore) error {
	errEmpty := validation.CheckDataEmpty(data.Name, data.Description, data.Location, data.MaxMembers)
	if errEmpty != nil {
		return errEmpty
	}

	// Mengubah huruf pertama di Name menjadi huruf besar
	titleCase := cases.Title(language.Indonesian)
	data.Name = titleCase.String(data.Name)

	// Mengubah huruf pertama di Location menjadi huruf besar
	data.Location = titleCase.String(data.Location)

	_, err := cs.communityRepository.GetByName(data.Name)
	if err == nil {
		return errors.New("nama community sudah digunakan")
	}

	errCreate := cs.communityRepository.CreateCommunity(image, data)
	if errCreate != nil {
		return errCreate
	}
	return nil
}

// DeleteCommunityById implements entity.CommunityServiceInterface.
func (cs *communityService) DeleteCommunityById(id string) error {
	if id == "" {
		return errors.New(constanta.ERROR_ID_INVALID)
	}

	err := cs.communityRepository.DeleteCommunityById(id)
	if err != nil {
		return err
	}

	return nil
}

// GetAllCommunity implements entity.CommunityServiceInterface.
func (cs *communityService) GetAllCommunity(page, limit, search string) ([]entity.CommunityCore, pagination.PageInfo, int, error) {
	pageInt, limitInt, err := validation.ValidateTypePaginationParameter(limit, page)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	pageValid, limitValid := validation.ValidateCountLimitAndPage(pageInt, limitInt)

	dropPointCores, pageInfo, count, err := cs.communityRepository.GetAllCommunity(pageValid, limitValid, search)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	return dropPointCores, pageInfo, count, nil
}

// GetCommunityById implements entity.CommunityServiceInterface.
func (cs *communityService) GetCommunityById(id string) (entity.CommunityCore, error) {
	if id == "" {
		return entity.CommunityCore{}, errors.New(constanta.ERROR_ID_INVALID)
	}

	idCommunity, err := cs.communityRepository.GetCommunityById(id)
	if err != nil {
		return entity.CommunityCore{}, err
	}

	return idCommunity, err
}

// UpdateCommunityById implements entity.CommunityServiceInterface.
func (cs *communityService) UpdateCommunityById(id string, image *multipart.FileHeader, data entity.CommunityCore) error {
	errEmpty := validation.CheckDataEmpty(data.Name, data.Description, data.Location, data.MaxMembers)
	if errEmpty != nil {
		return errEmpty
	}

	titleCase := cases.Title(language.Indonesian)
	data.Name = titleCase.String(data.Name)
	data.Location = titleCase.String(data.Location)

	_, err := cs.communityRepository.GetByName(data.Name)
	if err == nil {
		return errors.New("nama community sudah digunakan")
	}

	err = cs.communityRepository.UpdateCommunityById(id, image, data)
	if err != nil {
		return err
	}

	return nil
}
