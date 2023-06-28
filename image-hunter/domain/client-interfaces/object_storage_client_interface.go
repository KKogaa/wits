package clientinterfaces

import "github.com/KKogaa/image-hunter/domain/entities"

type ObjectStorageClient interface {
	SaveObject(*entities.Content) (*entities.Content, error)
	GetObject(string) (*entities.Content, error)
}
