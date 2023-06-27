package presenters

import "github.com/KKogaa/image-hunter/domain/entities"

type VectorPresenter struct {
	Vector     []float32
	Algorithm  string
	Path       string
	Similarity float32
}

func NewVectorPresenter(vector *entities.Vector) *VectorPresenter {
	return &VectorPresenter{
		Vector:    vector.Vector,
		Algorithm: vector.Algorithm,
		Path:      vector.Path,
        Similarity:  vector.Similarity,
	}
}
