package clientinterfaces

import "github.com/KKogaa/image-hunter/domain/entities"

type HasherClient interface {
	GetVectorFromImage(imgUrl string) (*entities.Vector, error)
	GetVectorFromText(text string) (*entities.Vector, error)
}
