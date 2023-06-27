package clientinterfaces

import "github.com/KKogaa/image-hunter/domain/entities"

type VectorStorageClient interface {
	GetSimilarVector(vector *entities.Vector) ([]*entities.Vector, error)
	SaveVector(vector *entities.Vector) (*entities.Vector, error)
}
