package wishlist

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (c *connector) Select(userid string, wishlistid string) (*models.Wishlist, error) {

	bwWishlist := BWWishlist{}

	sql := fmt.Sprintf("SELECT _id, meta_user_id, spec_name, spec_description FROM %s WHERE _id = '%s' AND meta_user_id = '%s'",
		c.tablename,
		wishlistid, userid)
	c.logging.Info(sql)
	err := c.db.QueryRow(sql).
		Scan(&bwWishlist.ID, &bwWishlist.MetaUserID, &bwWishlist.SpecName, &bwWishlist.SpecDescription)
	if err != nil {
		return nil, err
	}

	return fromWishlistDbModelToWishlistModel(bwWishlist), nil
}
