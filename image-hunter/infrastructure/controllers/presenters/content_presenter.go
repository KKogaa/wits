package presenters

import "github.com/KKogaa/image-hunter/domain/entities"


type ContentPresenter struct {
	Name    string
	Path    string
	Vectors []VectorPresenter
}

func NewContentPresenter(content *entities.Content) ContentPresenter {

    var vectorsPresenter []VectorPresenter
    for _, vector := range content.Vectors {
        vectorsPresenter = append(vectorsPresenter, VectorPresenter{
            Vector: vector.Vector,
            Path: vector.Path,
            Algorithm: vector.Algorithm,
        })


    }
    return ContentPresenter{
        Name: content.Name, 
        Path: content.Path,
        Vectors: vectorsPresenter,
    }
}
