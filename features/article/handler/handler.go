package handler

import (
	"net/http"
	"recything/features/article/dto/request"
	"recything/features/article/dto/response"
	"recything/features/article/entity"
	"recything/utils/helper"
	"recything/utils/jwt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type articleHandler struct {
	articleService entity.ArticleServiceInterface
}

func NewArticleHandler(article entity.ArticleServiceInterface) *articleHandler {
	return &articleHandler{
		articleService: article,
	}
}

func (article *articleHandler) CreateArticle(e echo.Context) error {
	Id, role, _ := jwt.ExtractToken(e)
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan id"))
	}
	if role == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan role"))
	}

	if role != "admin" && role != "super_admin"{
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("akses ditolak"))
	}

	newArticle := request.ArticleRequest{}
	err := e.Bind(&newArticle)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	image, err := e.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse("tidak ada file yang di upload"))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal upload file"))
	}

	articleInput := request.ArticleRequestToArticleCore(newArticle)
	_, errCreate := article.articleService.CreateArticle(articleInput, image)
	if errCreate != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse("berhasil menambahkan artikel"))
}

func (article *articleHandler) GetAllArticle(e echo.Context) error {
	tittle := e.QueryParam("tittle")
	page, _ := strconv.Atoi(e.QueryParam("page"))
	limit, _ := strconv.Atoi(e.QueryParam("limit"))

	articleData, paginationInfo, err := article.articleService.GetAllArticle(page, limit, tittle)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan artikel"))
	}

	var articleResponse = response.ListArticleCoreToListArticleResponse(articleData)

	return e.JSON(http.StatusOK, helper.SuccessWithPagnation("berhasil mendapatkan semua article", articleResponse, paginationInfo))
}

func (article *articleHandler) GetSpecificArticle(e echo.Context) error {
	idParams := e.Param("id")

	articleData, err := article.articleService.GetSpecificArticle(idParams)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse("gagal membaca data"))
	}

	var articleResponse = response.ArticleCoreToArticleResponse(articleData)

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("berhasil mendapatkan artikel", articleResponse))
}

func (article *articleHandler) UpdateArticle(e echo.Context) error {
	Id, role, _ := jwt.ExtractToken(e)
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan id"))
	}
	if role == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan role"))
	}

	if role != "admin" && role != "super_admin"{
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("akses ditolak"))
	}

	idParams := e.Param("id")

	updatedData := request.ArticleRequest{}
	errBind := e.Bind(&updatedData)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errBind.Error()))
	}

	image, err := e.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse("tidak ada file yang di upload"))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal upload file"))
	}

	articleInput := request.ArticleRequestToArticleCore(updatedData)
	updateArticle, errCreate := article.articleService.UpdateArticle(idParams, articleInput, image)
	if errCreate != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(errCreate.Error()))
	}

	articleResponse := response.ArticleCoreToArticleResponse(updateArticle)
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("berhasil", articleResponse))
}

func (article *articleHandler) DeleteArticle(e echo.Context) error {
	Id, role, _ := jwt.ExtractToken(e)
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan id"))
	}
	if role == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan role"))
	}

	if role != "admin" && role != "super_admin"{
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("akses ditolak"))
	}

	idParams := e.Param("id")

	errDelete := article.articleService.DeleteArticle(idParams)
	if errDelete != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(errDelete.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil menghapus artikel"))
}
