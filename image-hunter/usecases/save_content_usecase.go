package usecases

import (
	clientinterfaces "github.com/KKogaa/image-hunter/domain/client-interfaces"
	"github.com/KKogaa/image-hunter/domain/entities"
)

type SaveContentUsecase struct {
	hasherClient     clientinterfaces.HasherClient
	osClient         clientinterfaces.ObjectStorageClient
	vsClient clientinterfaces.VectorStorageClient
}

func NewSaveContentUsecase(hasherClient clientinterfaces.HasherClient,
	osClient clientinterfaces.ObjectStorageClient,
	vsClient clientinterfaces.VectorStorageClient) *SaveContentUsecase {
	return &SaveContentUsecase{
		hasherClient:     hasherClient,
		osClient:         osClient,
		vsClient: vsClient,
	}
}

func (s SaveContentUsecase) SaveImage(image []byte,
	imageName string) (*entities.Content, error) {

	savedContent, err := s.osClient.SaveObject(&entities.Content{
		Name: imageName,
		Data: image,
	})
	if err != nil {
		return nil, err
	}

	vector, err := s.hasherClient.GetVectorFromImage(savedContent.Path)
	if err != nil {
		return nil, err
	}
	vector.Path = savedContent.Path

	if _, err := s.vsClient.SaveVector(vector); err != nil {
		return nil, err
	}

	savedContent.AppendVector(vector)

	return savedContent, nil
}

func (s SaveContentUsecase) SaveImageUrl(imageUrl string,
	imageName string) (*entities.Content, error) {

	vector, err := s.hasherClient.GetVectorFromImage(imageUrl)
	if err != nil {
		return nil, err
	}

	vector.Path = imageUrl 

	if _, err := s.vsClient.SaveVector(vector); err != nil {
		return nil, err
	}

    savedContent := &entities.Content{
        Name: imageName,
        Path: imageUrl,
        Vectors: []*entities.Vector{vector},
    }

	return savedContent, nil
}
