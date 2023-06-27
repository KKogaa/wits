package usecases

import (
	clientinterfaces "github.com/KKogaa/image-hunter/domain/client-interfaces"
	"github.com/KKogaa/image-hunter/domain/entities"
)

type SearchImageUsecase struct {
	hasherClient     clientinterfaces.HasherClient
	vectorStorageClient clientinterfaces.VectorStorageClient
}

func NewSearchImageUsecase(hasherClient clientinterfaces.HasherClient,
	vectorStorageClient clientinterfaces.VectorStorageClient) *SearchImageUsecase {
	return &SearchImageUsecase{
		hasherClient:     hasherClient,
		vectorStorageClient: vectorStorageClient,
	}
}

func (s SearchImageUsecase) SearchVectorsSimilarToText(textImage string) ([]*entities.Vector, error) {

	vector, err := s.hasherClient.GetVectorFromText(textImage)

	if err != nil {
		return nil, err
	}

	similarVectors, err := s.vectorStorageClient.GetSimilarVector(vector)

	if err != nil {
		return nil, err
	}

	return similarVectors, nil
}

func (s SearchImageUsecase) SearchVectorsSimilarToImage(imageUrl string) ([]*entities.Vector, error) {

	vector, err := s.hasherClient.GetVectorFromImage(imageUrl)

	if err != nil {
		return nil, err
	}

	similarVectors, err := s.vectorStorageClient.GetSimilarVector(vector)

	if err != nil {
		return nil, err
	}

	return similarVectors, nil
}
