package entities

type Vector struct {
	ID         string
	Vector     []float32
	Algorithm  string
	Path       string
	Similarity float32
}

func (v Vector) GetDimensions() int {
	return len(v.Vector)
}
