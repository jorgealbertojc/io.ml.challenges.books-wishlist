package wishlist

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (c *connector) Select(userid string, wishlistid string) (*models.Wishlist, error) {

	bwWishlist := BWWishlist{}

	sql := fmt.Sprintf("SELECT _id, meta_user_id, spec_name, spec_description FROM %s WHERE meta_user_id = '%s' AND _id = '%s'",
		c.tablename,
		userid, wishlistid)
	c.logging.Info(sql)
	err := c.db.QueryRow(sql).
		Scan(&bwWishlist.ID, &bwWishlist.MetaUserID, &bwWishlist.SpecName, &bwWishlist.SpecDescription)
	if err != nil {
		return nil, err
	}

	return fromWishlistDbModelToWishlistModel(bwWishlist), nil
}

func (c *connector) SelectByName(userid string, wishlistname string) (*models.Wishlist, error) {

	bwWishlist := BWWishlist{}

	sql := fmt.Sprintf("SELECT _id, meta_user_id, spec_name, spec_description FROM %s WHERE meta_user_id = '%s' AND spec_name = '%s'",
		c.tablename,
		userid, wishlistname)
	c.logging.Info(sql)
	err := c.db.QueryRow(sql).
		Scan(&bwWishlist.ID, &bwWishlist.MetaUserID, &bwWishlist.SpecName, &bwWishlist.SpecDescription)
	if err != nil {
		return nil, err
	}

	return fromWishlistDbModelToWishlistModel(bwWishlist), nil
}
