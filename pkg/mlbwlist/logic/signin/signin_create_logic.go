package signin

import (
	"crypto/md5"
	"fmt"

	"github.com/google/uuid"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (l *logic) Create(userAccountModel models.UserAccount) (*models.Signin, error) {

	signin := models.Signin{
		ID: uuid.New().String(),
		Spec: &models.SigninSpec{
			UserID:    userAccountModel.ID,
			TokenHash: fmt.Sprintf("%x", md5.Sum([]byte("example-token"))),
		},
	}

	err := l.db.Insert(signin)
	if err != nil {
		return nil, err
	}

	return &signin, nil
}
