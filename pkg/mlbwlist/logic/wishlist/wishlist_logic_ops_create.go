package wishlist

import (
	"fmt"

	"github.com/google/uuid"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (l *logic) Create(wishlistModel models.Wishlist) (*models.Wishlist, error) {

	wishlistModel.ID = uuid.New().String()

	err := l.validateParentEntity(wishlistModel)
	if err != nil {
		return nil, err
	}

	storedWishlistModel, _ := l.validateWishlistExistenceByName(wishlistModel)
	if storedWishlistModel != nil {
		return nil, fmt.Errorf("wishlist identified with {%s / %s} already exists in our records", wishlistModel.Meta.UserID, wishlistModel.Spec.Name)
	}

	err = l.db.Insert(wishlistModel)
	if err != nil {
		return nil, err
	}

	return &wishlistModel, nil
}
