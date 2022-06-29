package signin

import "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"

func fromSigninDBModelToSigninModel(bwUserAccountSigninToken BWUserAccountSigninToken) *models.Signin {

	signinModel := models.Signin{}

	signinModel.ID = bwUserAccountSigninToken.ID
	signinModel.Meta = &models.SigninMeta{}
	signinModel.Meta.UserID = bwUserAccountSigninToken.MetaUserID
	signinModel.Spec = &models.SigninSpec{}
	signinModel.Spec.TokenHash = bwUserAccountSigninToken.SpecTokenHash

	return &signinModel
}
