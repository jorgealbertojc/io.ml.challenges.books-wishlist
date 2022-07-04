package books

import "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"

func (l *logic) Read(userid string, wishlistid string, bookid string) (*models.Book, error) {

	bookModel := models.Book{Meta: &models.BookMeta{UserID: userid, WishlistID: wishlistid}}
	err := l.validateEntityParents(bookModel)
	if err != nil {
		return nil, err
	}

	book, err := l.db.Select(userid, wishlistid, bookid)
	if err != nil {
		return nil, err
	}

	return book, nil
}
