package clientinterfaces

import "github.com/KKogaa/image-hunter/domain/entities"

type VectorStorageClient interface {
	GetSimilarVector(embedding []float32) ([]*entities.Vector, error)
	SaveVector(imageId string, embedding []float32) (*entities.Vector, error)
	GetAllVectors() ([]*entities.Vector, error)
}
