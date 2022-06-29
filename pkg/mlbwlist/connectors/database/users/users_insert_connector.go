package users

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (c *connector) Insert(userAccountModel models.UserAccount) error {

	sql := fmt.Sprintf("INSERT INTO user_accounts(_id, created_at, spec_username, spec_password) VALUES('%s', %d, '%s', '%s')", userAccountModel.ID, userAccountModel.CreatedAt, userAccountModel.Spec.Username, userAccountModel.Spec.Password)
	c.logging.Info(sql)
	_, err := c.db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
