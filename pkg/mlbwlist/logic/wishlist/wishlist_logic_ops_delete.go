package wishlist

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (l *logic) Delete(userid string, wishlistid string) error {

	wishlistModel := models.Wishlist{ID: wishlistid, Meta: &models.WishlistMeta{UserID: userid}}
	err := l.validateParentEntity(wishlistModel)
	if err != nil {
		return err
	}

	storedWishlistModel, _ := l.validateWishlistExistence(wishlistModel)
	if storedWishlistModel == nil {
		return fmt.Errorf("wishlist identified with {%s / %s} as not found in our records", userid, wishlistid)
	}

	bookListModel, err := l.bookdb.List(userid, wishlistid)
	if err != nil {
		return err
	}

	if len(bookListModel.Items) > 0 {
		return fmt.Errorf("wishlist identified with {%s / %s} cannot be deleted because have %d registered books", userid, wishlistid, len(bookListModel.Items))
	}

	return l.db.Delete(userid, wishlistid)
}
