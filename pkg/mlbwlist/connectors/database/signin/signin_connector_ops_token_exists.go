package signin

import (
	"fmt"
)

func (c *connector) TokenExists(token string) error {

	signinTokenModel := BWUserAccountSigninToken{}
	sql := fmt.Sprintf("SELECT spec_token_hash FROM %s WHERE spec_token_hash = '%s'",
		c.tablename,
		token)
	c.logging.Info(sql)
	err := c.db.QueryRow(sql).
		Scan(&signinTokenModel.SpecTokenHash)
	if err != nil {
		return err
	}

	return nil
}
