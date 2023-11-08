package handler

import (
	"encoding/json"
	"net/http"
	"recything/features/user/dto/request"
	"recything/features/user/dto/response"
	"recything/features/user/entity"
	"recything/utils/email"
	"recything/utils/helper"
	"recything/utils/jwt"
	"strings"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUseCase entity.UsersUsecaseInterface
}

func NewUserHandlers(uc entity.UsersUsecaseInterface) *userHandler {
	return &userHandler{
		userUseCase: uc,
	}
}

func (uh *userHandler) ForgetPassword(c echo.Context) error {
	newPassword := request.UserForgetPassword{}
	decoder := json.NewDecoder(c.Request().Body)
	decoder.DisallowUnknownFields()

	errBind := decoder.Decode(&newPassword)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid input"))
	}

	idToken, _, err := jwt.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}

	updateData := request.UsersRequestForgetPasswordToUsersCore(newPassword)
	_, err = uh.userUseCase.ForgetPassword(idToken, updateData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success update password"))
}

func (uh *userHandler) UpdateById(c echo.Context) error {
	dataUpdate := request.UserUpdate{}
	decoder := json.NewDecoder(c.Request().Body)
	decoder.DisallowUnknownFields()

	errBind := decoder.Decode(&dataUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid input"))
	}

	idToken, _, err := jwt.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}

	updateData := request.UsersRequestUpdateToUsersCore(dataUpdate)

	_, err = uh.userUseCase.UpdateById(idToken, updateData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data"))

}

func (uh *userHandler) Register(c echo.Context) error {
	// Bind data
	dataInput := request.UserRegister{}
	decoder := json.NewDecoder(c.Request().Body)
	decoder.DisallowUnknownFields()

	errBind := decoder.Decode(&dataInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid input"))
	}

	data := request.UsersRequestRegisterToUsersCore(dataInput)

	errCreate := uh.userUseCase.Register(data)
	if errCreate != nil {
		if strings.Contains(errCreate.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse(errCreate.Error()))
		} else {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("failed to create data "+errCreate.Error()))
		}
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse("success create data"))
}

func (uh *userHandler) VerifyAccount(c echo.Context) error {
	token := c.QueryParam("token")

	alreadyVerified, err := uh.userUseCase.VerifyUser(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid or expired verification token"))
	}

	if alreadyVerified {
		emailDone, err := email.ParseTemplate("verification_active.html", nil)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to parse template"))
		}
		return c.HTML(http.StatusOK, emailDone)
	}

	emailContent, err := email.ParseTemplate("success_verification.html", nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to parse template"))
	}
	return c.HTML(http.StatusOK, emailContent)
}

func (uh *userHandler) Login(c echo.Context) error {
	// Bind data
	login := request.UserLogin{}
	errBind := c.Bind(&login)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("eror bind data"))
	}

	// Memanggil func di usecase
	user, token, err := uh.userUseCase.Login(login.Email, login.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("login failed"))
	}

	if !user.IsVerified {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("account not verified"))
	}
	jwt.SetTokenCookie(c, token)
	response := response.UsersCoreToLoginResponse(user)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("login successful", response))
}

func (uh *userHandler) GetUser(c echo.Context) error {
	// Extra token dari id
	idToken, _, err := jwt.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}

	result, err := uh.userUseCase.GetById(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error reading data"))
	}

	var usersResponse = response.UsersCoreToResponseProfile(result)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get profile", usersResponse))
}
