package wishlist

import "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"

type WishListDBModel struct {
	ID              string
	SpecUserID      string
	SpecName        string
	SpecDescription string
}

func fromWishlistDbModelToWishlistModel(wishlistDbModel WishListDBModel) *models.Wishlist {

	wishlistModel := models.Wishlist{}
	wishlistModel.ID = wishlistDbModel.ID
	wishlistModel.Spec = &models.WishlistSpec{}
	wishlistModel.Spec.User = wishlistDbModel.SpecUserID
	wishlistModel.Spec.Name = wishlistDbModel.SpecName
	wishlistModel.Spec.Description = wishlistDbModel.SpecDescription

	return &wishlistModel
}
