package users

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (c *connector) Select(userid string) (*models.UserAccount, error) {

	userAccount := BWUserAccount{}
	sql := fmt.Sprintf("SELECT _id, spec_username, spec_password FROM %s WHERE _id='%s'", c.tablename, userid)
	c.logging.Info(sql)
	err := c.db.QueryRow(sql).
		Scan(&userAccount.ID, &userAccount.SpecUsername, &userAccount.SpecPassword)
	if err != nil {
		return nil, err
	}

	return fromUserAccountToUserAccountModel(userAccount), nil
}

func (c *connector) SelectByUsername(username string) (*models.UserAccount, error) {

	bwUserAccount := BWUserAccount{}
	sql := fmt.Sprintf("SELECT _id, spec_username, spec_password FROM %s WHERE spec_username = '%s'",
		c.tablename,
		username)
	c.logging.Info(sql)
	err := c.db.QueryRow(sql).
		Scan(&bwUserAccount.ID, &bwUserAccount.SpecUsername, &bwUserAccount.SpecPassword)
	if err != nil {
		return nil, err
	}

	return fromUserAccountToUserAccountModel(bwUserAccount), nil
}
