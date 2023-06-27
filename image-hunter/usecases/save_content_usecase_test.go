package usecases_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/KKogaa/image-hunter/domain/entities"
	"github.com/KKogaa/image-hunter/usecases"
)

func TestSaveContentUsecase(t *testing.T) {
	t.Run("should return an error when s3 client fails", func(t *testing.T) {
		mockS3Client := MockObjectStorageClient{
			SaveObjectFunc: func(content *entities.Content) (*entities.Content, error) {
				return nil, errors.New("some error")
			},
		}

		mockHasherClient := MockHasherClient{}

		mockVectorRepository := MockVectorRepository{}

		usecase := usecases.NewSaveContentUsecase(mockHasherClient,
			mockS3Client, mockVectorRepository)

		image := []byte{1, 2, 3}
		imageName := "image-name"
		content, err := usecase.SaveImage(image, imageName)

		if content != nil {
			t.Errorf("expected content to be empty")
		}

		if err == nil {
			t.Errorf("expected error to occur")
		}
	})

	t.Run("should return an error when hasher client fails", func(t *testing.T) {
		image := []byte{1, 2, 3}
		imageName := "image-name"
		mockS3Client := MockObjectStorageClient{
			SaveObjectFunc: func(content *entities.Content) (*entities.Content, error) {
				return &entities.Content{
					ID:   "asdf",
					Name: "image-test.png",
					Path: "http://minio/image-test.png",
					Data: image,
				}, nil
			},
		}

		mockHasherClient := MockHasherClient{
			GetVectorFromImageFunc: func(imgUrl string) (*entities.Vector, error) {
				return nil, errors.New("some error")
			},
		}

		mockVectorRepository := MockVectorRepository{}

		usecase := usecases.NewSaveContentUsecase(mockHasherClient,
			mockS3Client, mockVectorRepository)

		content, err := usecase.SaveImage(image, imageName)

		if content != nil {
			t.Errorf("expected content to be empty")
		}

		if err == nil {
			t.Errorf("expected error to occur")
		}

	})

	t.Run("should return an error when vectorizer repo fails", func(t *testing.T) {

		imageName := "image-name"
		image := []byte{1, 2, 3}
		vector := []float32{0.32, 0.31, 0.51}

		mockS3Client := MockObjectStorageClient{
			SaveObjectFunc: func(content *entities.Content) (*entities.Content, error) {
				return &entities.Content{
					ID:   "asdf",
					Name: "image-test.png",
					Path: "http://minio/image-test.png",
					Data: image,
				}, nil
			},
		}

		mockHasherClient := MockHasherClient{
			GetVectorFromImageFunc: func(imgUrl string) (*entities.Vector, error) {
				return &entities.Vector{
					Vector:    vector,
					Algorithm: "clip",
				}, nil
			},
		}

		mockVectorRepository := MockVectorRepository{
			SaveVectorFunc: func(vector *entities.Vector) (*entities.Vector, error) {
				return nil, errors.New("some error")
			},
		}

		usecase := usecases.NewSaveContentUsecase(mockHasherClient,
			mockS3Client, mockVectorRepository)

		content, err := usecase.SaveImage(image, imageName)

		if content != nil {
			t.Errorf("expected content to be empty")
		}

		if err == nil {
			t.Errorf("expected error to occur")
		}

	})

	t.Run("should return a saved content", func(t *testing.T) {
		image := []byte{1, 2, 3}
		vector := entities.Vector{
			ID:        "asdljfkjsdlkfj",
			Vector:    []float32{0.32, 0.31, 0.52},
			Algorithm: "clip",
			Path:      "asdlfjadslkf",
		}
		content := entities.Content{
			ID:      "asdf",
			Name:    "image-test.png",
			Path:    "http://minio/image-test.png",
			Data:    image,
			Vectors: []*entities.Vector{&vector},
		}

		mockS3Client := MockObjectStorageClient{
			SaveObjectFunc: func(content *entities.Content) (*entities.Content, error) {
				return &entities.Content{
					ID:   "asdf",
					Name: "image-test.png",
					Path: "http://minio/image-test.png",
					Data: image,
				}, nil
			},
		}

		mockHasherClient := MockHasherClient{
			GetVectorFromImageFunc: func(imgUrl string) (*entities.Vector, error) {
				return &vector, nil
			},
		}

		mockVectorRepository := MockVectorRepository{
			SaveVectorFunc: func(vector *entities.Vector) (*entities.Vector, error) {
				return vector, nil
			},
		}

		usecase := usecases.NewSaveContentUsecase(mockHasherClient, mockS3Client, mockVectorRepository)

		savedContent, err := usecase.SaveImage(image, content.Path)

		if savedContent.ID != content.ID {
			t.Errorf("expected ID to be equal")
			t.Errorf("has %s, expected %s", savedContent.ID, content.ID)
		}
		if savedContent.Name != content.Name {
			t.Errorf("expected Name to be equal")
			t.Errorf("has %s, expected %s", savedContent.Name, content.Name)
		}
		if savedContent.Path != content.Path {
			t.Errorf("expected Path to be equal")
			t.Errorf("has %s, expected %s", savedContent.Path, content.Path)
		}
		if !reflect.DeepEqual(savedContent.Data, content.Data) {
			t.Errorf("expected Data to be equal")
		}

		if len(savedContent.Vectors) != len(content.Vectors) {
			t.Errorf("expected Vectors to have the same length")
		}

		for i := range savedContent.Vectors {
			if !reflect.DeepEqual(savedContent.Vectors[i], content.Vectors[i]) {
				t.Errorf("expected Vectors[%d] to be equal", i)
			}
		}

		if err != nil {
			t.Errorf("expected no error to occur")
		}
	})

}
