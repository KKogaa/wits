package usecases

import (
	clientinterfaces "github.com/KKogaa/image-hunter/domain/client-interfaces"
	"github.com/KKogaa/image-hunter/domain/entities"
)

type GetAllVectorsUsecase struct {
	vsClient clientinterfaces.VectorStorageClient
}

func NewGetAllVectorsUsecase(vsClient clientinterfaces.VectorStorageClient) *GetAllVectorsUsecase {
	return &GetAllVectorsUsecase{
		vsClient: vsClient,
	}
}

//TODO: reduce scope to filter by user
func (g GetAllVectorsUsecase) Execute() ([]*entities.Vector, error) {

	vectors, err := g.vsClient.GetAllVectors()

	if err != nil {
		return nil, err
	}

	return vectors, nil
}
