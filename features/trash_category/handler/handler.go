package handler

import (
	"fmt"
	"net/http"
	"recything/features/trash_category/dto/request"
	"recything/features/trash_category/dto/response"
	"recything/features/trash_category/entity"
	"recything/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type trashCategoryHandler struct {
	trashCategory entity.TrashCategoryServiceInterface
}

func NewTrashCategoryHandler(trashCategory entity.TrashCategoryServiceInterface) *trashCategoryHandler {
	return &trashCategoryHandler{trashCategory: trashCategory}
}

func (tc *trashCategoryHandler) CreateCategory(e echo.Context) error {
	requestCategory := request.TrashCategory{}
	err := helper.DecodeJSON(e, &requestCategory)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	input := request.RequestTrashCategoryToCoreTrashCategory(requestCategory)
	err = tc.trashCategory.CreateCategory(input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}
	return e.JSON(http.StatusCreated, helper.SuccessResponse("Berhasil menambahkan kategori sampah"))
}

func (tc *trashCategoryHandler) GetAllCategory(e echo.Context) error {
	page := e.QueryParam("page")
	limit := e.QueryParam("limit")
	trashType := e.QueryParam("trash_type")

	result, pagnation, err := tc.trashCategory.GetAllCategory(page, trashType, limit)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	pageInt, _ := strconv.Atoi(page)
	errMessage := fmt.Sprintf("page %s tidak ada", page)
	if pageInt > pagnation.LastPage {
		return e.JSON(http.StatusNotFound, helper.ErrorResponse(errMessage))
	}

	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse("Belum ada kategori sampah"))
	}

	response := response.ListCoreTrashCategoryToReponseTrashCategory(result)
	return e.JSON(http.StatusOK, helper.SuccessWithPagnationAndDataResponse("Berhasil mendapatkan seluruh kategori sampah", response, pagnation))
}

func (tc *trashCategoryHandler) GetById(e echo.Context) error {
	id := e.Param("id")
	
	result, err := tc.trashCategory.GetById(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	response := response.CoreTrashCategoryToReponseTrashCategory(result)
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
	requestCategory := request.TrashCategory{}
	err := helper.DecodeJSON(e, &requestCategory)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	input := request.RequestTrashCategoryToCoreTrashCategory(requestCategory)
	result, err := tc.trashCategory.UpdateCategory(id, input)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	response := response.CoreTrashCategoryToReponseTrashCategory(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Berhasil mengupdate kategori sampah", response))
}
