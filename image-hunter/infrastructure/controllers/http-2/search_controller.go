package http2

import (
	"log"
	"net/http"

	"github.com/KKogaa/image-hunter/infrastructure/controllers/http-2/dtos"
	"github.com/KKogaa/image-hunter/infrastructure/controllers/http-2/presenters"

	"github.com/KKogaa/image-hunter/usecases"
	"github.com/gin-gonic/gin"
)

type SearchController struct {
	searchImageUsecase *usecases.SearchImageUsecase
}

func NewSearchController(searchImageUsecase *usecases.SearchImageUsecase) *SearchController {
	return &SearchController{
		searchImageUsecase: searchImageUsecase,
	}
}

func (s SearchController) SetupRoutes(router *gin.Engine) {
	router.GET("/search/text", s.SearchImageWithText)
	router.GET("/search/image", s.SearchImageWithImageUrl)
}

func (s SearchController) SearchImageWithText(ctx *gin.Context) {

	var searchTextDTO dtos.SearchTextDTO
	if err := ctx.Bind(&searchTextDTO); err != nil {
		return
	}

	vectors, err := s.searchImageUsecase.SearchVectorsSimilarToText(searchTextDTO.Text)

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

func (s SearchController) SearchImageWithImageUrl(ctx *gin.Context) {

	var searchImageLinkDTO dtos.SearchImageLinkDTO
	if err := ctx.Bind(&searchImageLinkDTO); err != nil {
		return
	}

	vectors, err := s.searchImageUsecase.SearchVectorsSimilarToImage(searchImageLinkDTO.Url)

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
