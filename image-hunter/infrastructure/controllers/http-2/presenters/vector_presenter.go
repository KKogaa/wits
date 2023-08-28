package presenters

import "github.com/KKogaa/image-hunter/domain/entities"

type VectorPresenter struct {
	Embedding  []float32 `json:"embedding"`
	Path       string    `json:"path"`
	Similarity float32   `json:"similarity"`
}

func NewVectorPresenter(vector *entities.Vector) *VectorPresenter {
	return &VectorPresenter{
		Embedding:  vector.Embedding,
		Path:       vector.ImagePath,
		Similarity: vector.Similarity,
	}
}
