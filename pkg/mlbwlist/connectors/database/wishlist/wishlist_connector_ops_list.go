package wishlist

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (c *connector) List(userid string) (*models.WishlistList, error) {

	wishlistList := models.WishlistList{Items: []models.Wishlist{}}

	sql := fmt.Sprintf("SELECT _id, meta_user_id, spec_name, spec_description FROM %s WHERE meta_user_id = '%s'",
		c.tablename,
		userid)
	c.logging.Info(sql)
	rows, err := c.db.Query(sql)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		wbWishlist := BWWishlist{}
		err := rows.Scan(&wbWishlist.ID, &wbWishlist.MetaUserID, &wbWishlist.SpecName, &wbWishlist.SpecDescription)
		if err != nil {
			return nil, err
		}
		wishlistList.Items = append(wishlistList.Items, *fromWishlistDbModelToWishlistModel(wbWishlist))
	}

	return &wishlistList, nil
}
