package usecases

import (
	clientinterfaces "github.com/KKogaa/image-hunter/domain/client-interfaces"
	"github.com/KKogaa/image-hunter/domain/entities"
)

type SaveContentUsecase struct {
	hasherClient        clientinterfaces.HasherClient
	osClient            clientinterfaces.ObjectStorageClient
	vsClient            clientinterfaces.VectorStorageClient
	imageDownloadClient clientinterfaces.ImageDownloadClient
}

func NewSaveContentUsecase(hasherClient clientinterfaces.HasherClient,
	osClient clientinterfaces.ObjectStorageClient,
	vsClient clientinterfaces.VectorStorageClient,
	imageDownloadClient clientinterfaces.ImageDownloadClient,
) *SaveContentUsecase {
	return &SaveContentUsecase{
		hasherClient:        hasherClient,
		osClient:            osClient,
		vsClient:            vsClient,
		imageDownloadClient: imageDownloadClient,
	}
}

func (s SaveContentUsecase) SaveImage(image []byte,
	imageName string) (*entities.Vector, error) {

	imageId, err := s.osClient.SaveObject(image, imageName)
	if err != nil {
		return nil, err
	}

	imageUrl, err := s.osClient.GetObject(imageId)
	if err != nil {
		return nil, err
	}

	embedding, err := s.hasherClient.GetEmbeddingFromImage(imageUrl)
	if err != nil {
		return nil, err
	}

	vector, err := s.vsClient.SaveVector(imageId, embedding)
	if err != nil {
		return nil, err
	}

	vector.ImagePath = imageUrl

	return vector, nil
}

func (s SaveContentUsecase) SaveImageUrl(imageUrl string,
	imageName string) (*entities.Vector, error) {

	imageData, err := s.imageDownloadClient.GetImage(imageUrl)
	if err != nil {
		return nil, err
	}

	imageId, err := s.osClient.SaveObject(imageData, imageName)
	if err != nil {
		return nil, err
	}

	privateImageUrl, err := s.osClient.GetObject(imageId)
	if err != nil {
		return nil, err
	}

	embedding, err := s.hasherClient.GetEmbeddingFromImage(privateImageUrl)
	if err != nil {
		return nil, err
	}

	vector, err := s.vsClient.SaveVector(imageId, embedding)
	if err != nil {
		return nil, err
	}

	vector.ImagePath = imageUrl

	return vector, nil
}
