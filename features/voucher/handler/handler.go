package handler

import (
	"net/http"
	"recything/features/voucher/dto/request"
	"recything/features/voucher/dto/response"
	"recything/features/voucher/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
)

type voucherHandler struct {
	VoucherService entity.VoucherServiceInterface
}

func NewVoucherHandler(voucher entity.VoucherServiceInterface) *voucherHandler {
	return &voucherHandler{VoucherService: voucher}
}

func (vh *voucherHandler) CreateVoucher(e echo.Context) error {
	input := request.VoucherRequest{}
	_, role, errExtract := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	err := helper.BindFormData(e, &input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	image, err := e.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ERROR_EMPTY_FILE))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal upload file"))
	}

	request := request.RequestVoucherToCoreVoucher(input)
	err = vh.VoucherService.Create(image, request)
	if err != nil {
		if helper.HttpResponseCondition(err, constanta.ERROR_MESSAGE...) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse("Berhasil menambahkan data"))
}

func (vh *voucherHandler) GetAllVoucher(e echo.Context) error {
	_, _, errExtract := jwt.ExtractToken(e)

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	page := e.QueryParam("page")
	limit := e.QueryParam("limit")
	search := e.QueryParam("search")

	result, pagination, count, err := vh.VoucherService.GetAll(page, limit, search)
	if err != nil {
		if helper.HttpResponseCondition(err, constanta.ERROR_MESSAGE...) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_NULL))
	}

	response := response.ListCoreVoucherToCoreVoucher(result)
	return e.JSON(http.StatusOK, helper.SuccessWithPagnationAndCount("Berhasil mendapatkan seluruh data", response, pagination, count))
}

func (vh *voucherHandler) GetVoucherById(e echo.Context) error {
	_, _, errExtract := jwt.ExtractToken(e)

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	id := e.Param("id")
	result, err := vh.VoucherService.GetById(id)
	if err != nil {
		if helper.HttpResponseCondition(err, constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))

		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	response := response.CoreVoucherToResponVoucher(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Berhasil mendapatkan data", response))
}

func (vh *voucherHandler) UpdateVoucher(e echo.Context) error {
	input := request.VoucherRequest{}

	id := e.Param("id")

	_, role, errExtract := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	err := e.Bind(&input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	image, err := e.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ERROR_EMPTY_FILE))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	request := request.RequestVoucherToCoreVoucher(input)
	err = vh.VoucherService.UpdateData(id, image, request)
	if err != nil {
		if helper.HttpResponseCondition(err, constanta.ERROR_DATA_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(err.Error()))
		}
		if helper.HttpResponseCondition(err, constanta.ERROR_MESSAGE...) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}

		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Berhasil mengupdate data"))
}

func (vh *voucherHandler) DeleteVoucherById(e echo.Context) error {
	_, role, errExtract := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	id := e.Param("id")
	err := vh.VoucherService.DeleteData(id)
	if err != nil {
		if helper.HttpResponseCondition(err, constanta.ERROR_DATA_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Berhasil menghapus data"))
}
