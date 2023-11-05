package handler

import (
	"net/http"
	"recything/features/admin/dto"
	"recything/features/admin/entity"
	"recything/utils/helper"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	AdminService entity.AdminServiceInterface
}

func NewAdminHandler(admin entity.AdminServiceInterface) *AdminHandler {
	return &AdminHandler{AdminService: admin}
}

func (admin *AdminHandler) Create(e echo.Context) error {
	_, role := jwt.ExtractToken(e)
	if role != helper.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse("failed"))
	}

	inputAdmin := dto.AdminRequest{}

	if err := e.Bind(&inputAdmin); err != nil {
		return err
	}

	adminCore := entity.AdminRequestToAdminCore(inputAdmin)
	adminCreated, err := admin.AdminService.Create(adminCore)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	adminResponse := entity.AdminCoreToAdminResponse(adminCreated)
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("succes", adminResponse))

}

func (admin *AdminHandler) GetAll(e echo.Context) error {
	_, role := jwt.ExtractToken(e)
	if role != helper.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse("failed"))
	}

	AdminsData, err := admin.AdminService.GetAll()
	if err != nil {
		e.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed"))
	}

	adminsResponse := entity.ListAdminCoreToAdminResponse(AdminsData)
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("succes", adminsResponse))

}
func (admin *AdminHandler) GetById(e echo.Context) error {
	adminId := e.Param("id")

	_, role := jwt.ExtractToken(e)
	if role != helper.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse("failed"))
	}

	AdminData, err := admin.AdminService.GetById(adminId)
	if err != nil {
		e.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed"))
	}

	adminResponse := entity.AdminCoreToAdminResponse(AdminData)
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("succes", adminResponse))
}

func (admin *AdminHandler) Delete(e echo.Context) error {
	adminId := e.Param("id")

	_, role := jwt.ExtractToken(e)
	if role != helper.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse("failed"))
	}

	err := admin.AdminService.DeleteById(adminId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("success"))
}

func (admin *AdminHandler) UpdateById(e echo.Context) error {
	adminId := e.Param("id")

	_, role := jwt.ExtractToken(e)
	if role != helper.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse("failed"))
	}

	newAdmin := dto.AdminRequest{}
	err := e.Bind(&newAdmin)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("failed"))
	}

	coreAdmin := entity.AdminRequestToAdminCore(newAdmin)
	err = admin.AdminService.UpdateById(adminId, coreAdmin)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}
	return e.JSON(http.StatusOK, helper.SuccessResponse("success"))
}

func (admin *AdminHandler) Login(e echo.Context) error {
	input := dto.AdminRequest{}
	if err := e.Bind(&input); err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	adminData, token, err := admin.AdminService.FindByEmailANDPassword(input.Email, input.Password)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	jwt.SetTokenCookie(e, token)
	adminResponse := entity.AdminCoreToAdminResponse(adminData)

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Succes Login", adminResponse))
}
