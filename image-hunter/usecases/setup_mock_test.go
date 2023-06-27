package usecases_test

import "github.com/KKogaa/image-hunter/domain/entities"

type MockVectorRepository struct {
	GetSimilarVectorFunc func(vector *entities.Vector) ([]*entities.Vector, error)
	SaveVectorFunc       func(vector *entities.Vector) (*entities.Vector, error)
}

func (m MockVectorRepository) GetSimilarVector(vector *entities.Vector) ([]*entities.Vector, error) {
	return m.GetSimilarVectorFunc(vector)
}

func (m MockVectorRepository) SaveVector(vector *entities.Vector) (*entities.Vector, error) {
	return m.SaveVectorFunc(vector)
}

type MockHasherClient struct {
	GetVectorFromImageFunc func(imgUrl string) (*entities.Vector, error)
	GetVectorFromTextFunc  func(text string) (*entities.Vector, error)
}

func (m MockHasherClient) GetVectorFromImage(imgUrl string) (*entities.Vector, error) {
	return m.GetVectorFromImageFunc(imgUrl)
}

func (m MockHasherClient) GetVectorFromText(text string) (*entities.Vector, error) {
	return m.GetVectorFromTextFunc(text)
}

type MockObjectStorageClient struct {
	SaveObjectFunc func(content *entities.Content) (*entities.Content, error)
}

func (m MockObjectStorageClient) SaveObject(content *entities.Content) (*entities.Content, error) {
	return m.SaveObjectFunc(content)
}
