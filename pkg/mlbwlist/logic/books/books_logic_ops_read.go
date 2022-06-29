package books

import "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"

func (l *logic) Read(userid string, wishlistid string, bookid string) (*models.Book, error) {

	book, err := l.db.Select(userid, wishlistid, bookid)
	if err != nil {
		return nil, err
	}

	return book, nil
}
