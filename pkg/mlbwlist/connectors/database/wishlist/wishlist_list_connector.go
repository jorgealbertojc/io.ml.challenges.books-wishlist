package wishlist

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (c *connector) List(userid string) (*models.WishlistList, error) {

	wishlistList := models.WishlistList{Items: []models.Wishlist{}}

	sql := fmt.Sprintf("SELECT _id, spec_user_id, spec_name, spec_description FROM books_wishlists WHERE spec_user_id = '%s'", userid)
	c.logging.Info(sql)
	rows, err := c.db.Query(sql)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		wlistDbModel := WishListDBModel{}
		err := rows.Scan(&wlistDbModel.ID, &wlistDbModel.SpecUserID, &wlistDbModel.SpecName, &wlistDbModel.SpecDescription)
		if err != nil {
			return nil, err
		}
		wishlistList.Items = append(wishlistList.Items, *fromWishlistDbModelToWishlistModel(wlistDbModel))
	}

	return &wishlistList, nil
}
