package wishlist

import "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"

func (l *logic) List(userid string) (*models.WishlistList, error) {

	list, err := l.db.List(userid)
	if err != nil {
		return nil, err
	}

	return list, nil
}
