package entities

type Vector struct {
	ID         uint64
	Embedding  []float32
	Algorithm  string
	Similarity float32
	ImageID    string
	ImagePath  string
}

func (v Vector) GetDimensions() int {
	return len(v.Embedding)
}
