package usecases_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/KKogaa/image-hunter/domain/entities"
	"github.com/KKogaa/image-hunter/usecases"
)

func TestSearchVectors(t *testing.T) {
	t.Run("searchbytext usecase should return an error when hasher client fails", func(t *testing.T) {

		text := "this is a test text"

		mockHasherClient := MockHasherClient{
			GetVectorFromTextFunc: func(text string) (*entities.Vector, error) {
				return nil, errors.New("some error")
			},
		}

		mockVectorRepository := MockVectorRepository{}

		usecase := usecases.NewSearchImageUsecase(mockHasherClient,
			mockVectorRepository)

		vectors, err := usecase.SearchVectorsSimilarToText(text)

		if err == nil {
			t.Errorf("expected error to occur")
		}

		if vectors != nil {
			t.Errorf("expected to return nil")
		}
	})

	t.Run("searchbytext usecase should return an error when vector repository fails", func(t *testing.T) {

		text := "this is a test text"

		mockHasherClient := MockHasherClient{
			GetVectorFromTextFunc: func(text string) (*entities.Vector, error) {
				return &entities.Vector{
					Vector:    []float32{0.32, 0.56, 0.75},
					Algorithm: "clip",
				}, nil
			},
		}

		mockVectorRepository := MockVectorRepository{
			GetSimilarVectorFunc: func(vector *entities.Vector) ([]*entities.Vector, error) {
				return nil, errors.New("some error")
			},
		}

		usecase := usecases.NewSearchImageUsecase(mockHasherClient,
			mockVectorRepository)

		vectors, err := usecase.SearchVectorsSimilarToText(text)

		if err == nil {
			t.Errorf("expected error to occur")
		}

		if vectors != nil {
			t.Errorf("expected to return nil")
		}
	})

	t.Run("searchbytext usecase should return similar vectors to text", func(t *testing.T) {

		vectors := []*entities.Vector{
			{
				ID:        "1",
				Vector:    []float32{0.33, 0.44, 0.45},
				Algorithm: "clip",
				Path:      "aklsdjf;lkajsdflk",
			},
		}

		text := "this is a test text"

		mockHasherClient := MockHasherClient{
			GetVectorFromTextFunc: func(text string) (*entities.Vector, error) {
				return &entities.Vector{
					Vector:    []float32{0.32, 0.56, 0.75},
					Algorithm: "clip",
				}, nil
			},
		}

		mockVectorRepository := MockVectorRepository{
			GetSimilarVectorFunc: func(vector *entities.Vector) ([]*entities.Vector, error) {
				return vectors, nil

			},
		}

		usecase := usecases.NewSearchImageUsecase(mockHasherClient,
			mockVectorRepository)

		searchedVectors, err := usecase.SearchVectorsSimilarToText(text)

		if err != nil {
			t.Errorf("expected no error to occur")
		}

		if !reflect.DeepEqual(searchedVectors, vectors) {
			t.Errorf("expected vectors to be the same")
		}

	})

	t.Run("searchbyimage usecase should return an error when hasher client fails", func(t *testing.T) {

		imageUrl := "sample-url.com"

		mockHasherClient := MockHasherClient{
			GetVectorFromImageFunc: func(text string) (*entities.Vector, error) {
				return nil, errors.New("some error")
			},
		}

		mockVectorRepository := MockVectorRepository{}

		usecase := usecases.NewSearchImageUsecase(mockHasherClient,
			mockVectorRepository)

		vectors, err := usecase.SearchVectorsSimilarToImage(imageUrl)

		if err == nil {
			t.Errorf("expected error to occur")
		}

		if vectors != nil {
			t.Errorf("expected to return nil")
		}
	})

	t.Run("searchbyimage usecase should return an error when vector repository fails", func(t *testing.T) {

		imageUrl := "sample-url.com"

		mockHasherClient := MockHasherClient{
			GetVectorFromImageFunc: func(text string) (*entities.Vector, error) {
				return &entities.Vector{
					Vector:    []float32{0.32, 0.56, 0.75},
					Algorithm: "clip",
				}, nil
			},
		}

		mockVectorRepository := MockVectorRepository{
			GetSimilarVectorFunc: func(vector *entities.Vector) ([]*entities.Vector, error) {
				return nil, errors.New("some error")
			},
		}

		usecase := usecases.NewSearchImageUsecase(mockHasherClient,
			mockVectorRepository)

		vectors, err := usecase.SearchVectorsSimilarToImage(imageUrl)

		if err == nil {
			t.Errorf("expected error to occur")
		}

		if vectors != nil {
			t.Errorf("expected to return nil")
		}
	})

	t.Run("searchbyimage usecase should return similar vectors to image url", func(t *testing.T) {

		vectors := []*entities.Vector{
			{
				ID:        "1",
				Vector:    []float32{0.33, 0.44, 0.45},
				Algorithm: "clip",
				Path:      "aklsdjf;lkajsdflk",
			},
		}

		mockHasherClient := MockHasherClient{
			GetVectorFromImageFunc: func(text string) (*entities.Vector, error) {
				return &entities.Vector{
					Vector:    []float32{0.32, 0.56, 0.75},
					Algorithm: "clip",
				}, nil
			},
		}

		mockVectorRepository := MockVectorRepository{
			GetSimilarVectorFunc: func(vector *entities.Vector) ([]*entities.Vector, error) {
				return vectors, nil

			},
		}

		imageUrl := "sample-url.com"

		usecase := usecases.NewSearchImageUsecase(mockHasherClient,
			mockVectorRepository)

		searchedVectors, err := usecase.SearchVectorsSimilarToImage(imageUrl)

		if err != nil {
			t.Errorf("expected no error to occur")
		}

		if !reflect.DeepEqual(searchedVectors, vectors) {
			t.Errorf("expected vectors to be the same")
		}

	})

}
