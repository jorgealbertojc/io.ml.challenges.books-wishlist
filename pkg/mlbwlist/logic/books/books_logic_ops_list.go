package books

import (
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (l *logic) List(userid string, wishlistid string) (*models.BookList, error) {

	list, err := l.db.List(userid, wishlistid)
	if err != nil {
		return nil, err
	}

	return list, nil
}
