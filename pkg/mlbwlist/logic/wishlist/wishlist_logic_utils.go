package wishlist

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (l *logic) validateParentEntity(wishlistModel models.Wishlist) error {

	userAccountModel, _ := l.userdb.Select(wishlistModel.Meta.UserID)
	if userAccountModel == nil {
		return fmt.Errorf("user {%s} cannot be found in our records", wishlistModel.Meta.UserID)
	}

	return nil
}

func (l *logic) validateWishlistExistenceByName(wishlistModel models.Wishlist) (*models.Wishlist, error) {

	storedWishlistModel, err := l.db.SelectByName(wishlistModel.Meta.UserID, wishlistModel.Spec.Name)
	if err != nil {
		return nil, err
	}

	return storedWishlistModel, nil
}

func (l *logic) validateWishlistExistence(wishlistModel models.Wishlist) (*models.Wishlist, error) {

	storedWishlistModel, err := l.db.Select(wishlistModel.Meta.UserID, wishlistModel.ID)
	if err != nil {
		return nil, err
	}

	if storedWishlistModel == nil {
		return nil, fmt.Errorf("wishlist identified with {%s / %s} was not found in our records", wishlistModel.Meta.UserID, wishlistModel.ID)
	}

	return storedWishlistModel, nil
}
