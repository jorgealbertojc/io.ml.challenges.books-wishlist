package signin

import (
	"fmt"

	"github.com/google/uuid"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (l *logic) Create(userAccountModel models.UserAccount) (*models.Signin, error) {

	storedUserAccountModel, _ := l.userdb.SelectByUsername(userAccountModel.Spec.Username)
	if storedUserAccountModel == nil {
		return nil, fmt.Errorf("user account {%s} was not found in our records", userAccountModel.Spec.Username)
	}

	signinToken, err := l.buildSigninAuthenticationToken()
	if err != nil {
		return nil, err
	}

	signinModel := models.Signin{
		ID: uuid.New().String(),
		Meta: &models.SigninMeta{
			UserID: storedUserAccountModel.ID,
		},
		Spec: &models.SigninSpec{
			TokenHash: signinToken,
		},
	}

	err = l.db.Insert(signinModel)
	if err != nil {
		return nil, err
	}

	return &signinModel, nil
}
