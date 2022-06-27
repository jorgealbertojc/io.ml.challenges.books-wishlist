package signin

import (
	"encoding/json"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (l *logic) Create(userAccountModel models.UserAccount) error {

	userAccountModelJSON, _ := json.MarshalIndent(userAccountModel, "", "    ")
	l.logging.Info("%s", string(userAccountModelJSON))

	return nil
}
