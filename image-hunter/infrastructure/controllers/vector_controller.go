package controllers

import (
	"log"
	"net/http"

	"github.com/KKogaa/image-hunter/infrastructure/controllers/presenters"

	"github.com/KKogaa/image-hunter/usecases"
	"github.com/gin-gonic/gin"
)

type VectorController struct {
	getAllVectorsUsecase *usecases.GetAllVectorsUsecase
}

func NewVectorController(getAllVectors *usecases.GetAllVectorsUsecase) *VectorController {
	return &VectorController{
		getAllVectorsUsecase: getAllVectors,
	}
}

func (v VectorController) SetupRoutes(router *gin.Engine) {
	router.GET("/vectors", v.GetAll)
}

func (v VectorController) GetAll(ctx *gin.Context) {

	vectors, err := v.getAllVectorsUsecase.Execute()

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	vectorsPresenter := make([]presenters.VectorPresenter, len(vectors))
	for i, vector := range vectors {
		vectorsPresenter[i] = *presenters.NewVectorPresenter(vector)
	}
	ctx.JSON(http.StatusOK, vectorsPresenter)
}
