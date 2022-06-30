package users

import (
	"github.com/google/uuid"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (l *logic) Create(user models.UserAccount) (*models.UserAccount, error) {

	user.ID = uuid.New().String()

	err := l.db.Insert(user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
