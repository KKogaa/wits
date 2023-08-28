package clientinterfaces

type ObjectStorageClient interface {
	SaveObject(data []byte, name string) (string, error)
	GetObject(imageId string) (string, error)
}
