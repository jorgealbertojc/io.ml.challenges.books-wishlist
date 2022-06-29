package users

import "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"

type UserDBModel struct {
	ID           string
	CreatedAt    int
	SpecUsername string
	SpecPassword string
}

func fromUserDBModelToUserAccountModel(userDbModel UserDBModel) *models.UserAccount {

	userAccount := models.UserAccount{
		ID:        userDbModel.ID,
		CreatedAt: userDbModel.CreatedAt,
		Spec: &models.UserAccountSpec{
			Username: userDbModel.SpecUsername,
			Password: userDbModel.SpecPassword,
		},
	}

	return &userAccount
}

func fromUserAccountModelToUserDbModel(userAccountModel *models.UserAccount) UserDBModel {

	userDBModel := UserDBModel{}
	userDBModel.ID = userAccountModel.ID
	userDBModel.CreatedAt = userAccountModel.CreatedAt
	if userAccountModel.Spec != nil {
		userDBModel.SpecUsername = userAccountModel.Spec.Username
		userDBModel.SpecPassword = userAccountModel.Spec.Password
	}

	return userDBModel
}
