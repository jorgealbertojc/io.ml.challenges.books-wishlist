package search

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (l *logic) Find(queryargs map[string][]string) (*models.BookList, error) {

	bookList := models.BookList{Items: []models.Book{}}
	queryparams := map[string]string{}
	q, exists := queryargs["q"]
	if exists {
		queryparams["q"] = q[0]
	}

	if len(queryparams["q"]) == 0 {
		return nil, fmt.Errorf("cannot execute search without ")
	}

	searchResultList, err := l.searchConnector.Find(queryparams)
	if err != nil {
		return nil, err
	}

	for _, book := range searchResultList.Items {

		bookItem := models.Book{}
		bookItem.Meta = &models.BookMeta{}
		bookItem.Meta.GoogleID = book.ID
		bookItem.Spec = &models.BookSpec{}
		bookItem.Spec.Title = book.VolumeInfo.Title
		bookItem.Spec.Publisher = book.VolumeInfo.Publisher
		bookItem.Spec.Authors = book.VolumeInfo.Authors
		bookList.Items = append(bookList.Items, bookItem)
	}

	return &bookList, nil
}
