package clientinterfaces

type ImageDownloadClient interface {
	GetImage(imageUrl string) ([]byte, error)
}
