package books

import (
	"github.com/google/uuid"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (l *logic) Create(bookModel models.Book) (*models.Book, error) {

	bookModel.ID = uuid.New().String()

	result, err := l.searchConnector.Read(bookModel.Meta.GoogleID)
	if err != nil {
		return nil, err
	}

	bookModel.Spec.Title = result.VolumeInfo.Title
	bookModel.Spec.Authors = result.VolumeInfo.Authors
	bookModel.Spec.Publisher = result.VolumeInfo.Publisher

	err = l.db.Insert(bookModel)
	if err != nil {
		return nil, err
	}

	return &bookModel, nil
}
