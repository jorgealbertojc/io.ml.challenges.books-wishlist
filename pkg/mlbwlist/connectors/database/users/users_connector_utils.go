package users

import "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"

func fromUserAccountToUserAccountModel(bwUserAccount BWUserAccount) *models.UserAccount {

	userAccountModel := models.UserAccount{
		ID: bwUserAccount.ID,
		Spec: &models.UserAccountSpec{
			Username: bwUserAccount.SpecUsername,
			Password: bwUserAccount.SpecPassword,
		},
	}

	return &userAccountModel
}
