package users

import "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"

func (l *logic) Read(userid string) (*models.UserAccount, error) {

	user := models.UserAccount{
		ID: userid,
	}
	return &user, nil
}
