package usecases

import (
	clientinterfaces "github.com/KKogaa/image-hunter/domain/client-interfaces"
	"github.com/KKogaa/image-hunter/domain/entities"
)

type SearchImageUsecase struct {
	hasherClient        clientinterfaces.HasherClient
	vectorStorageClient clientinterfaces.VectorStorageClient
	osClient            clientinterfaces.ObjectStorageClient
}

func NewSearchImageUsecase(hasherClient clientinterfaces.HasherClient,
	vectorStorageClient clientinterfaces.VectorStorageClient,
	osClient clientinterfaces.ObjectStorageClient,

) *SearchImageUsecase {
	return &SearchImageUsecase{
		hasherClient:        hasherClient,
		vectorStorageClient: vectorStorageClient,
		osClient:            osClient,
	}
}

func (s SearchImageUsecase) SearchVectorsSimilarToText(textImage string) ([]*entities.Vector, error) {

	embedding, err := s.hasherClient.GetEmbeddingFromText(textImage)

	if err != nil {
		return nil, err
	}

	similarVectors, err := s.vectorStorageClient.GetSimilarVector(embedding)

	if err != nil {
		return nil, err
	}

	for _, vector := range similarVectors {
		imageUrl, err := s.osClient.GetObject(vector.ImageID)
		if err != nil {
			return nil, err
		}

		vector.ImagePath = imageUrl
	}

	return similarVectors, nil
}

func (s SearchImageUsecase) SearchVectorsSimilarToImage(imageUrl string) ([]*entities.Vector, error) {

	embedding, err := s.hasherClient.GetEmbeddingFromImage(imageUrl)

	if err != nil {
		return nil, err
	}

	similarVectors, err := s.vectorStorageClient.GetSimilarVector(embedding)

	if err != nil {
		return nil, err
	}

	for _, vector := range similarVectors {
		imageUrl, err := s.osClient.GetObject(vector.ImageID)
		if err != nil {
			return nil, err
		}

		vector.ImagePath = imageUrl
	}

	return similarVectors, nil

}
