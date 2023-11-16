package handler

import (
	"net/http"
	req "recything/features/trash_category/dto/request"
	resp "recything/features/trash_category/dto/response"
	"recything/features/trash_category/entity"
	"recything/utils/helper"

	"github.com/labstack/echo/v4"
)

type trashCategoryHandler struct {
	trashCategory entity.TrashCategoryServiceInterface
}

func NewTrashCategoryHandler(trashCategory entity.TrashCategoryServiceInterface) *trashCategoryHandler {
	return &trashCategoryHandler{trashCategory: trashCategory}
}

func (tc *trashCategoryHandler) CreateCategory(e echo.Context) error {
	request := req.TrashCategory{}
	err := helper.DecodeJSON(e, &request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	input := req.RequestTrashCategoryToCoreTrashCategory(request)
	result, err := tc.trashCategory.CreateCategory(input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	response := resp.CoreTrashCategoryToReponseTrashCategory(result)
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("Berhasil menambahkan kategori sampah", response))
}

func (tc *trashCategoryHandler) GetAllCategory(e echo.Context) error {
	page := e.QueryParam("page")
	limit := e.QueryParam("limit")
	result, pagnation, err := tc.trashCategory.GetAllCategory(page, limit)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	response := resp.ListCoreTrashCategoryToReponseTrashCategory(result)
	// return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Berhasil mendapatkan seluruh kategori sampah", response), )
	return e.JSON(200, echo.Map{
		"data":       response,
		"paganation": pagnation})
}

func (tc *trashCategoryHandler) GetById(e echo.Context) error {
	id := e.Param("id")
	result, err := tc.trashCategory.GetById(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	response := resp.CoreTrashCategoryToReponseTrashCategory(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Berhasil mendapatkan detail kategori sampah", response))
}

func (tc *trashCategoryHandler) DeleteById(e echo.Context) error {
	id := e.Param("id")
	err := tc.trashCategory.DeleteCategory(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Berhasil menghapus kategori"))
}

func (tc *trashCategoryHandler) UpdateCategory(e echo.Context) error {
	id := e.Param("id")
	request := req.TrashCategory{}
	err := helper.DecodeJSON(e, &request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	input := req.RequestTrashCategoryToCoreTrashCategory(request)
	result, err := tc.trashCategory.UpdateCategory(id, input)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	response := resp.CoreTrashCategoryToReponseTrashCategory(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Berhasil mengupdate kategori sampah", response))
}
