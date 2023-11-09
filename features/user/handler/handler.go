package handler

import (
	"encoding/json"
	"log"
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

func (uh *userHandler) UpdatePassword(c echo.Context) error {
	newPassword := request.UserUpdatePassword{}
	decoder := json.NewDecoder(c.Request().Body)
	decoder.DisallowUnknownFields()

	errBind := decoder.Decode(&newPassword)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("input salah"))
	}

	idToken, _, err := jwt.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}

	updateData := request.UsersRequestUpdatePasswordToUsersCore(newPassword)
	_, err = uh.userUseCase.UpdatePassword(idToken, updateData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("berhasil update password"))
}

func (uh *userHandler) ForgetPassword(c echo.Context) error {
	newPassword := request.UserForgetPassword{}
	decoder := json.NewDecoder(c.Request().Body)
	decoder.DisallowUnknownFields()

	errBind := decoder.Decode(&newPassword)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("input salah"))
	}

	email, err := jwt.ExtractTokenVerifikasi(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}
	log.Println("emailnya :", email)
	updateData := request.UsersRequestForgetPasswordToUsersCore(newPassword)
	err = uh.userUseCase.ForgetPassword(email, updateData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("berhasil update password"))
}

func (uh *userHandler) UpdateById(c echo.Context) error {
	dataUpdate := request.UserUpdate{}
	decoder := json.NewDecoder(c.Request().Body)
	decoder.DisallowUnknownFields()

	errBind := decoder.Decode(&dataUpdate)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("input salah"))
	}

	idToken, _, err := jwt.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("failed to extra token " + err.Error()))
	}

	updateData := request.UsersRequestUpdateToUsersCore(dataUpdate)

	_, err = uh.userUseCase.UpdateById(idToken, updateData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("failed to update data " + err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("berhasil update data"))

}

func (uh *userHandler) Register(c echo.Context) error {
	// Bind data
	dataInput := request.UserRegister{}
	decoder := json.NewDecoder(c.Request().Body)
	decoder.DisallowUnknownFields()

	errBind := decoder.Decode(&dataInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("input salah"))
	}

	data := request.UsersRequestRegisterToUsersCore(dataInput)

	errCreate := uh.userUseCase.Register(data)
	if errCreate != nil {
		if strings.Contains(errCreate.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse(errCreate.Error()))
		} else {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal membuat data "+errCreate.Error()))
		}
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse("berhasil create data"))
}

func (uh *userHandler) VerifyAccount(c echo.Context) error {
	token := c.QueryParam("token")

	alreadyVerified, err := uh.userUseCase.VerifyUser(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("token telah kadaluarsa atau salah"))
	}

	if alreadyVerified {
		emailDone, err := email.ParseTemplate("verification_active.html", nil)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("gagal menguraikan template"))
		}
		return c.HTML(http.StatusOK, emailDone)
	}

	emailContent, err := email.ParseTemplate("success_verification.html", nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("gagal menguraikan template"))
	}
	return c.HTML(http.StatusOK, emailContent)
}

func (uh *userHandler) Login(c echo.Context) error {
	// Bind data
	login := request.UserLogin{}
	errBind := c.Bind(&login)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("gagal mengambil data"))
	}

	// Memanggil func di usecase
	user, token, err := uh.userUseCase.Login(login.Email, login.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("login gagal"))
	}

	if !user.IsVerified {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("akun tidak terverifikasi"))
	}
	jwt.SetTokenCookie(c, token)
	response := response.UsersCoreToLoginResponse(user)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("login berhasil", response))
}

func (uh *userHandler) GetUser(c echo.Context) error {
	// Extra token dari id
	idToken, _, err := jwt.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}

	result, err := uh.userUseCase.GetById(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("gagal membaca data"))
	}

	var usersResponse = response.UsersCoreToResponseProfile(result)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil mendapatkan profile", usersResponse))
}

func (uh *userHandler) EmailOTP(c echo.Context) error {
	dataInput := request.UserSendOTP{}
	errBind := c.Bind(&dataInput)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("gagal mengambil data"))
	}

	userCore := request.UsersRequestOTPToUsersCore(dataInput)

	err := uh.userUseCase.SendOTP(userCore.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("gagal mengirim OTP"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("otp berhasil dikirim"))
}

func (uh *userHandler) VerifyOTP(c echo.Context) error {
	dataInput := request.UserVerifyOTP{}
	errBind := c.Bind(&dataInput)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("gagal mengambil data"))
	}

	userCore := request.UsersRequestVerifyOTPToUsersCore(dataInput)

	token, err := uh.userUseCase.VerifyOTP(userCore.Email, userCore.Otp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("gagal verifikasi OTP " + err.Error()))
	}
	jwt.SetTokenCookie(c, token)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("verifikasi OTP berhasil", token))
}
