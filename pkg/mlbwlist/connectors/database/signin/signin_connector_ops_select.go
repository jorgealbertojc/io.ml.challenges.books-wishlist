package signin

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (c *connector) Select(userid string, token string) (*models.Signin, error) {

	bwUserAccountSigninToken := BWUserAccountSigninToken{}

	sql := fmt.Sprintf("SELECT _id, spec_user_id, spec_token_hash FROM %s WHERE spec_user_id = '%s' AND spec_token_hash = '%s'",
		c.tablename,
		userid, token)
	c.logging.Info(sql)
	err := c.db.QueryRow(sql).
		Scan(&bwUserAccountSigninToken.ID, &bwUserAccountSigninToken.MetaUserID, &bwUserAccountSigninToken.SpecTokenHash)
	if err != nil {
		return nil, err
	}

	return fromSigninDBModelToSigninModel(bwUserAccountSigninToken), nil
}
