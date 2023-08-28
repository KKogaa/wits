package clientinterfaces

type HasherClient interface {
	GetEmbeddingFromImage(imageUrl string) ([]float32, error)
	GetEmbeddingFromText(text string) ([]float32, error)
}
