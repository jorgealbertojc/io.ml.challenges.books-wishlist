package books

import "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"

func (l *logic) Read(userid string, wishlistid string, bookid string) (*models.Books, error) {

	book := models.Books{}
	book.ID = bookid
	return &book, nil
}
