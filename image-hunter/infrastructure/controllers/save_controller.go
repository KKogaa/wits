package controllers

import (
	"bytes"
	"io"
	"net/http"

	"github.com/KKogaa/image-hunter/infrastructure/controllers/dtos"
	"github.com/KKogaa/image-hunter/infrastructure/controllers/presenters"
	"github.com/KKogaa/image-hunter/usecases"
	"github.com/gin-gonic/gin"
)

type SaveController struct {
	saveUsecase *usecases.SaveContentUsecase
}

func NewSaveController(saveUsecase *usecases.SaveContentUsecase) *SaveController {
	return &SaveController{
		saveUsecase: saveUsecase,
	}
}

func (s SaveController) SetupRoutes(router *gin.Engine) {
	router.POST("/upload/file", s.SaveContentFromImage)
	router.POST("/upload/url", s.SaveContentFromImageUrl)
}

func (s SaveController) SaveContentFromImage(ctx *gin.Context) {
	formFile, header, err := ctx.Request.FormFile("file")

	defer formFile.Close()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	buffer := bytes.NewBuffer(nil)
	if _, err := io.Copy(buffer, formFile); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	content, err := s.saveUsecase.SaveImage(buffer.Bytes(), header.Filename)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, presenters.NewContentPresenter(content))
}

func (s SaveController) SaveContentFromImageUrl(ctx *gin.Context) {

	var saveContentDTO dtos.SaveContentDTO
	if err := ctx.ShouldBindJSON(&saveContentDTO); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	content, err := s.saveUsecase.SaveImageUrl(saveContentDTO.Url,
		saveContentDTO.Name)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, presenters.NewContentPresenter(content))
}
