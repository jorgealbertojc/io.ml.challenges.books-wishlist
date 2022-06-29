package wishlist

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (c *connector) Select(userid string, wishlistid string) (*models.Wishlist, error) {

	wlistDbModel := WishListDBModel{}

	sql := fmt.Sprintf("SELECT _id, spec_user_id, spec_name, spec_description FROM books_wishlists WHERE _id = '%s' AND spec_user_id = '%s'", wishlistid, userid)
	c.logging.Info(sql)
	err := c.db.QueryRow(sql).
		Scan(&wlistDbModel.ID, &wlistDbModel.SpecUserID, &wlistDbModel.SpecName, &wlistDbModel.SpecDescription)
	if err != nil {
		return nil, err
	}

	return fromWishlistDbModelToWishlistModel(wlistDbModel), nil
}
