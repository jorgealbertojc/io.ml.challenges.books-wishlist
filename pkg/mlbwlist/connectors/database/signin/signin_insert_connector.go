package signin

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (c *connector) Insert(signinModel models.Signin) error {

	sql := fmt.Sprintf("INSERT INTO users_signing_tokens(_id, spec_user_id, spec_token_hash) VALUES('%s', '%s', '%s')", signinModel.ID, signinModel.Spec.UserID, signinModel.Spec.TokenHash)
	c.logging.Info(sql)
	_, err := c.db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
