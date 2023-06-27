package entities

type Content struct {
	ID      string
	Name    string
	Path    string
	Data    []byte
	Vectors []*Vector
}

func (c *Content) AppendVector(vector *Vector) {
	c.Vectors = append(c.Vectors, vector)
}
