package handler

import (
	"net/http"
	"recything/features/admin/dto"
	"recything/features/admin/entity"
	"recything/utils/helper"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	AdminService entity.AdminServiceInterface
}

func NewAdminHandler(admin entity.AdminServiceInterface) *AdminHandler {
	return &AdminHandler{AdminService: admin}
}

func (admin *AdminHandler) Create(e echo.Context) {
	inputAdmin := dto.AdminRequest{}
	adminCore := entity.AdminRequestToAdminCore(inputAdmin)

	adminCreated, err := admin.AdminService.Create(adminCore)
	if err != nil {
		e.JSON(http.StatusInternalServerError, helper.FailedResponse("failed"))
	}

	adminResp := entity.AdminCoreToAdminResponse(adminCreated)
	e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("succes", adminResp))

}

func (admin *AdminHandler) GetAll(e echo.Context) {
	AdminsData, err := admin.AdminService.GetAll()
	

	
}
