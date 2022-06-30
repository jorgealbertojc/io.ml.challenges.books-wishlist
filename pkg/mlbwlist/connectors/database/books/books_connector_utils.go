package books

import (
	"encoding/json"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func fromBWWishlistBookToBookModel(bwWishlistBook BWWishlistBook) *models.Book {

	bookModel := models.Book{}

	authors := []string{}
	_ = json.Unmarshal([]byte(bwWishlistBook.SpecAuthors), &authors)

	bookModel.ID = bwWishlistBook.ID
	bookModel.Meta = &models.BookMeta{}
	bookModel.Meta.UserID = bwWishlistBook.MetaUserID
	bookModel.Meta.WishlistID = bwWishlistBook.MetaWishlistID
	bookModel.Meta.GoogleID = bwWishlistBook.MetaGoogleID
	bookModel.Spec = &models.BookSpec{}
	bookModel.Spec.Title = bwWishlistBook.SpecTitle
	bookModel.Spec.Authors = authors
	bookModel.Spec.Publisher = bwWishlistBook.SpecPublisher

	return &bookModel
}
