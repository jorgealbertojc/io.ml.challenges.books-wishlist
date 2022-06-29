package signin

import "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"

type SigninDBModel struct {
	ID            string
	SpecUserID    string
	SpecTokenHash string
}

func fromSigninDBModelToSigninModel(signinDbModel SigninDBModel) *models.Signin {

	signinModel := models.Signin{}

	signinModel.ID = signinDbModel.ID
	signinModel.Spec = &models.SigninSpec{}
	signinModel.Spec.UserID = signinDbModel.SpecUserID
	signinModel.Spec.TokenHash = signinDbModel.SpecTokenHash

	return &signinModel
}
