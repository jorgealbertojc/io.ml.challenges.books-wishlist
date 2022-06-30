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
