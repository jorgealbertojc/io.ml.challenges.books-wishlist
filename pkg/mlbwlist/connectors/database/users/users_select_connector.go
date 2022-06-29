package users

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (c *connector) Select(userid string) (*models.UserAccount, error) {

	dbmodel := UserDBModel{}
	sql := fmt.Sprintf("SELECT _id, created_at, spec_username, spec_password FROM user_accounts WHERE _id='%s'", userid)
	c.logging.Info(sql)
	err := c.db.QueryRow(sql).
		Scan(&dbmodel.ID, &dbmodel.CreatedAt, &dbmodel.SpecUsername, &dbmodel.SpecPassword)
	if err != nil {
		return nil, err
	}

	return fromUserDBModelToUserAccountModel(dbmodel), nil
}
