package wishlist

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (c *connector) Insert(wishlistModel models.Wishlist) error {

	sql := fmt.Sprintf("INSERT INTO %s(_id, meta_user_id, spec_name, spec_description) VALUES('%s', '%s', '%s', '%s')",
		c.tablename,
		wishlistModel.ID, wishlistModel.Meta.UserID, wishlistModel.Spec.Name, wishlistModel.Spec.Description)
	c.logging.Info(sql)
	_, err := c.db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
