package wishlist

import (
	"github.com/google/uuid"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (l *logic) Create(wishlistModel models.Wishlist) (*models.Wishlist, error) {

	wishlistModel.ID = uuid.New().String()

	err := l.db.Insert(wishlistModel)
	if err != nil {
		return nil, err
	}

	return &wishlistModel, nil
}