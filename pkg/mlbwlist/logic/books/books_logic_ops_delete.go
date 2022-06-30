package books

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (l *logic) Delete(userid string, wishlistid string, bookid string) error {

	bookModel := models.Book{ID: bookid, Meta: &models.BookMeta{UserID: userid, WishlistID: wishlistid}}
	err := l.validateEntityParents(bookModel)
	if err != nil {
		return err
	}

	storedBookModel, _ := l.validateBookExistence(bookModel)
	if storedBookModel == nil {
		return fmt.Errorf("book identified by {%s / %s / %s} was not found in or records", userid, wishlistid, bookid)
	}

	return l.db.Delete(userid, wishlistid, bookid)
}
