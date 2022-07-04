package users

import (
	"fmt"

	"github.com/google/uuid"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (l *logic) Create(userAccountModel models.UserAccount) (*models.UserAccount, error) {

	userAccountModel.ID = uuid.New().String()

	storedUserAccount, _ := l.db.SelectByUsername(userAccountModel.Spec.Username)
	if storedUserAccount != nil && storedUserAccount.Spec.Username == userAccountModel.Spec.Username {
		l.logging.Error("seems that user %s is already registered in our records", userAccountModel.Spec.Username)
		return nil, fmt.Errorf("user account username {%s} already exists in records", userAccountModel.Spec.Username)
	}

	err := l.db.Insert(userAccountModel)
	if err != nil {
		l.logging.Error("")
		return nil, err
	}

	return &userAccountModel, nil
}
