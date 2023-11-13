package handler

import (
	"net/http"
	"recything/features/admin/dto/request"
	"recything/features/admin/dto/response"
	"recything/features/admin/entity"
	user "recything/features/user/entity"
	userDto "recything/features/user/dto/response"
	reportRequest "recything/features/report/dto/request"
	reportDto "recything/features/report/dto/response"
	"recything/utils/helper"
	"recything/utils/jwt"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	AdminService entity.AdminServiceInterface
	UserService  user.UsersUsecaseInterface
}

func NewAdminHandler(as entity.AdminServiceInterface, us user.UsersUsecaseInterface) *AdminHandler {
	return &AdminHandler{
		AdminService: as,
		UserService:  us,
	}
}

// membuat admin, hanya untuk super admin
func (ah *AdminHandler) Create(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)

	if role != helper.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse("Acces Denied"))
	}
	if err != nil {
		return err
	}

	input := request.AdminRequest{}

	err = helper.DecodeJSON(e, &input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	request := request.AdminRequestToAdminCore(input)

	result, err := ah.AdminService.Create(request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	response := response.AdminCoreToAdminResponse(result)
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("berhasil membuat data admin", response))

}

// login untuk admin dan juga super admin
func (ah *AdminHandler) Login(e echo.Context) error {
	input := request.AdminLogin{}

	err := helper.DecodeJSON(e, &input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	request := request.RequestLoginToAdminCore(input)

	result, token, err := ah.AdminService.FindByEmailANDPassword(request)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	jwt.SetTokenCookie(e, token)
	response := response.AdminCoreToAdminResponse(result)

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil melakukan login", response))
}

// mendapatkan semua data admin yang active maupun yang tidak active
func (ah *AdminHandler) GetAll(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if role != helper.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse("Acces Denied"))
	}
	if err != nil {
		return err
	}

	result, err := ah.AdminService.GetAll()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))	
	}

	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse("data admin belum ada"))	
	}

	response := response.ListAdminCoreToAdminResponse(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil mengambil semua data admin", response))

}

// mendapatkan data admin detail lengkap
func (ah *AdminHandler) GetById(e echo.Context) error {
	
	adminId:= e.Param("id")

	_, role, err := jwt.ExtractToken(e)
	if role != helper.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse("Acces Denied"))
	}
	if err != nil {
		return err
	}

	result, err := ah.AdminService.GetById(adminId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	// if len(result.Id) == 0 {
	// 	return e.JSON(http.StatusOK, helper.SuccessResponse("data admin belum ada"))	
	// }

	response := response.AdminCoreToAdminResponse(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil mengambil data admin", response))
}

// menghapus data admin 
func (ah *AdminHandler) Delete(e echo.Context) error {
	adminId := e.Param("id")

	_, role, err := jwt.ExtractToken(e)
	if role != helper.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse("Acces Denied"))
	}
	if err != nil {
		return err
	}

	err = ah.AdminService.DeleteById(adminId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil menghapus data admin"))
}

// melakukan pembaruan atau edit data admin
func (ah *AdminHandler) UpdateById(e echo.Context) error {
	adminId := e.Param("id")

	_, role, err := jwt.ExtractToken(e)
	if role != helper.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse("Acces Denied"))
	}
	if err != nil {
		return err
	}

	input := request.AdminRequest{}

	err = helper.DecodeJSON(e, &input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	request := request.AdminRequestToAdminCore(input)
	err = ah.AdminService.UpdateById(adminId, request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}
	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil melakukan pembaruan data admin"))
}


// Manage User
func (ah *AdminHandler) GetAllUser(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if role != helper.SUPERADMIN && role != helper.ADMIN {
        return e.JSON(http.StatusForbidden, helper.ErrorResponse("unauthorized"))
    }

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse("failed extra token"))
	}

	result, err := ah.AdminService.GetAllUsers()
	if err != nil {
		e.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed"))
	}

	response := userDto.UsersCoreToResponseUsersList(result)
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("succes", response))

}

func (ah *AdminHandler) GetByIdUsers(e echo.Context) error {
	userId := e.Param("id")

	_, role, err := jwt.ExtractToken(e)
	if role != helper.SUPERADMIN && role != helper.ADMIN {
        return e.JSON(http.StatusForbidden, helper.ErrorResponse("unauthorized"))
    }

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse("failed extra token"))
	}

	UsersData, err := ah.AdminService.GetByIdUsers(userId)
	if err != nil {
		e.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed"))
	}

	userResponse := userDto.UsersCoreToResponseUsers(UsersData)
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("succes", userResponse))
}

func (ah *AdminHandler) DeleteUsers(e echo.Context) error {
	userId := e.Param("id")

	_, role, err := jwt.ExtractToken(e)
	if role != helper.SUPERADMIN && role != helper.ADMIN {
        return e.JSON(http.StatusForbidden, helper.ErrorResponse("unauthorized"))
    }

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse("failed extra token"))
	}

	err = ah.AdminService.DeleteUsers(userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("success"))
}

// Manage Reporting
func (ah *AdminHandler) GetByStatusReport(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)

	if role != helper.SUPERADMIN && role != helper.ADMIN {
        return e.JSON(http.StatusForbidden, helper.ErrorResponse("unauthorized"))
    }

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse("failed extra token"))
	}

	status := e.QueryParam("status")
	result, err := ah.AdminService.GetByStatusReport(status)
	if err != nil {
		if err.Error() == "status tidak valid" {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse("input status salah"))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse("gagal mendapatkan data"))
	}

	response := reportDto.ListReportCoresToReportResponseForDataReporting(result, ah.UserService)
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("success", response))

}

func (ah *AdminHandler) UpdateStatusReport(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)

	if role != helper.SUPERADMIN && role != helper.ADMIN {
        return e.JSON(http.StatusForbidden, helper.ErrorResponse("unauthorized"))
    }

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse("failed extra token"))
	}

	id := e.Param("id")

	input := reportRequest.UpdateStatusReportRubbish{}
	err = helper.DecodeJSON(e, &input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}
	
	_, err = ah.AdminService.UpdateStatusReport(id, input.Status)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil memperbarui status"))
}