package books

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (l *logic) validateEntityParents(bookModel models.Book) error {

	userAccountModel, _ := l.userdb.Select(bookModel.Meta.UserID)
	if userAccountModel == nil {
		return fmt.Errorf("user {%s} cannot be found in our records", bookModel.Meta.UserID)
	}

	wishlistModel, _ := l.wishlistdb.Select(bookModel.Meta.UserID, bookModel.Meta.WishlistID)
	if wishlistModel == nil {
		return fmt.Errorf("wishlist {%s / %s} cannot be found in our records", bookModel.Meta.UserID, bookModel.Meta.WishlistID)
	}

	return nil
}

func (l *logic) validateBookExistenceByGID(bookModel models.Book) (*models.Book, error) {

	storedBookModel, _ := l.db.SelectByGID(bookModel.Meta.UserID, bookModel.Meta.WishlistID, bookModel.Meta.GoogleID)
	if storedBookModel != nil {
		return storedBookModel, nil
	}

	return nil, fmt.Errorf("book identified with {%s / %s / gid: %s} already exists in records", bookModel.Meta.UserID, bookModel.Meta.WishlistID, bookModel.Meta.GoogleID)
}

func (l *logic) validateBookExistence(bookModel models.Book) (*models.Book, error) {

	storedBookModel, _ := l.db.Select(bookModel.Meta.UserID, bookModel.Meta.WishlistID, bookModel.ID)
	if storedBookModel != nil {
		return storedBookModel, nil
	}

	return nil, fmt.Errorf("book identified with {%s / %s / %s} already exists in records", bookModel.Meta.UserID, bookModel.Meta.WishlistID, bookModel.ID)
}
