package books

import (
	"github.com/google/uuid"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (l *logic) Create(userid string, wislistid string, bookModel models.Books) (*models.Books, error) {

	bookModel.ID = uuid.New().String()

	return &bookModel, nil
}
