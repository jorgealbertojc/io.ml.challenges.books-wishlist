package signin

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (c *connector) Select(userid string, token string) (*models.Signin, error) {

	signinDbModel := SigninDBModel{}

	sql := fmt.Sprintf("SELECT _id, spec_user_id, spec_token_hash FROM users_signing_tokens WHERE spec_user_id = '%s' AND spec_token_hash = '%s'", userid, token)
	c.logging.Info(sql)
	err := c.db.QueryRow(sql).
		Scan(&signinDbModel.ID, &signinDbModel.SpecUserID, &signinDbModel.SpecTokenHash)
	if err != nil {
		return nil, err
	}

	return fromSigninDBModelToSigninModel(signinDbModel), nil
}
