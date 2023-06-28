package presenters

import "github.com/KKogaa/image-hunter/domain/entities"

type VectorPresenter struct {
	Vector     []float32 `json:"vector"`
	Algorithm  string    `json:"algorithm"`
	Path       string    `json:"path"`
	Similarity float32   `json:"similarity"`
}

func NewVectorPresenter(vector *entities.Vector) *VectorPresenter {
	return &VectorPresenter{
		Vector:     vector.Vector,
		Algorithm:  vector.Algorithm,
		Path:       vector.Path,
		Similarity: vector.Similarity,
	}
}
