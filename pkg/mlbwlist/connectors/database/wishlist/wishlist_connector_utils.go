package wishlist

import "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"

func fromWishlistDbModelToWishlistModel(wishlistDbModel BWWishlist) *models.Wishlist {

	wishlistModel := models.Wishlist{}
	wishlistModel.ID = wishlistDbModel.ID
	wishlistModel.Spec = &models.WishlistSpec{}
	wishlistModel.Spec.User = wishlistDbModel.MetaUserID
	wishlistModel.Spec.Name = wishlistDbModel.SpecName
	wishlistModel.Spec.Description = wishlistDbModel.SpecDescription

	return &wishlistModel
}
