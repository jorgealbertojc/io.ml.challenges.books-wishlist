package signin

import (
	"crypto/md5"
	"encoding/json"
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (l *logic) Create(userAccountModel models.UserAccount) (*models.Signin, error) {

	userAccountModelJSON, _ := json.MarshalIndent(userAccountModel, "", "    ")
	l.logging.Info("%s", string(userAccountModelJSON))

	signin := models.Signin{
		Token: fmt.Sprintf("%x", md5.Sum([]byte("example-token"))),
	}

	return &signin, nil
}
