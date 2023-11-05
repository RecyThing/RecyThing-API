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
	_, role, err := jwt.ExtractToken(e)
	if role != helper.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse("Acces Denied"))
	}
	if err != nil {
		return err
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
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("create admin success", adminResponse))

}

func (admin *AdminHandler) GetAll(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if role != helper.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse("Acces Denied"))
	}
	if err != nil {
		return err
	}

	AdminsData, err := admin.AdminService.GetAll()
	if err != nil {
		e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	adminsResponse := entity.ListAdminCoreToAdminResponse(AdminsData)
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("all data admins", adminsResponse))

}
func (admin *AdminHandler) GetById(e echo.Context) error {
	adminId := e.Param("id")

	_, role, err := jwt.ExtractToken(e)
	if role != helper.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse("Acces Denied"))
	}
	if err != nil {
		return err
	}

	AdminData, err := admin.AdminService.GetById(adminId)
	if err != nil {
		e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	adminResponse := entity.AdminCoreToAdminResponse(AdminData)
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("data admin" + AdminData.Name, adminResponse))
}

func (admin *AdminHandler) Delete(e echo.Context) error {
	adminId := e.Param("id")

	_, role, err := jwt.ExtractToken(e)
	if role != helper.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse("Acces Denied"))
	}
	if err != nil {
		return err
	}

	err = admin.AdminService.DeleteById(adminId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("success delete admin"))
}

func (admin *AdminHandler) UpdateById(e echo.Context) error {
	adminId := e.Param("id")

	_, role, err := jwt.ExtractToken(e)
	if role != helper.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse("Acces Denied"))
	}
	if err != nil {
		return err
	}

	newAdmin := dto.AdminRequest{}
	err = e.Bind(&newAdmin)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	coreAdmin := entity.AdminRequestToAdminCore(newAdmin)
	err = admin.AdminService.UpdateById(adminId, coreAdmin)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}
	return e.JSON(http.StatusOK, helper.SuccessResponse("success update admin"))
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

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("success login", adminResponse))

}
