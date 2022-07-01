package search

import "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"

func fromSearchResultToBookItem(searchResultModel SearchResult) *models.Book {

	bookModel := models.Book{
		Meta: &models.BookMeta{},
		Spec: &models.BookSpec{},
	}

	bookModel.Meta.GoogleID = searchResultModel.ID
	bookModel.Spec.Title = searchResultModel.VolumeInfo.Title
	bookModel.Spec.Authors = searchResultModel.VolumeInfo.Authors
	bookModel.Spec.Publisher = searchResultModel.VolumeInfo.Publisher

	return &bookModel
}
