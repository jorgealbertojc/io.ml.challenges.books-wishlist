package wishlist

import "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"

func (l *logic) List(userid string) (*models.WishlistList, error) {

	wishlistModel := models.Wishlist{Meta: &models.WishlistMeta{UserID: userid}}
	err := l.validateParentEntity(wishlistModel)
	if err != nil {
		return nil, err
	}

	list, err := l.db.List(userid)
	if err != nil {
		return nil, err
	}

	return list, nil
}
