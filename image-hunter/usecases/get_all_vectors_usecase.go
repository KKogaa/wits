package usecases

import (
	clientinterfaces "github.com/KKogaa/image-hunter/domain/client-interfaces"
	"github.com/KKogaa/image-hunter/domain/entities"
)

type GetAllVectorsUsecase struct {
	vsClient clientinterfaces.VectorStorageClient
	osClient clientinterfaces.ObjectStorageClient
}

func NewGetAllVectorsUsecase(vsClient clientinterfaces.VectorStorageClient,
	osClient clientinterfaces.ObjectStorageClient) *GetAllVectorsUsecase {
	return &GetAllVectorsUsecase{
		vsClient: vsClient,
		osClient: osClient,
	}
}

//TODO: reduce scope to filter by user
func (g GetAllVectorsUsecase) Execute() ([]*entities.Vector, error) {

	vectors, err := g.vsClient.GetAllVectors()

	if err != nil {
		return nil, err
	}

	for _, vector := range vectors {
		imageUrl, err := g.osClient.GetObject(vector.ImageID)
		if err != nil {
			return nil, err
		}
		vector.ImagePath = imageUrl

	}

	return vectors, nil
}
