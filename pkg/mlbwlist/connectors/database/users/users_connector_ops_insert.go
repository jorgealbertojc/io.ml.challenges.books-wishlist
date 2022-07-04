package users

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (c *connector) Insert(userAccountModel models.UserAccount) error {

	sql := fmt.Sprintf("INSERT INTO %s(_id, spec_username, spec_password) VALUES('%s', '%s', '%s')",
		c.tablename,
		userAccountModel.ID, userAccountModel.Spec.Username, userAccountModel.Spec.Password)
	c.logging.Info(sql)
	_, err := c.db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
