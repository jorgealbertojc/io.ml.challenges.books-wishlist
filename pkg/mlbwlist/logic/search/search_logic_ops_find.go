package search

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (l *logic) Find(queryargs map[string][]string) (*models.BookList, error) {

	if len(queryargs) == 0 {
		return nil, fmt.Errorf("this action cannot be performed without query string arguments, allowed query string arguments 'q', 'title', 'author' or 'publisher'")
	}

	queryparams := map[string]string{}
	q, exists := queryargs["q"]
	if exists {
		queryparams["q"] = q[0]
	}

	intitle, exists := queryargs["title"]
	if exists {
		queryparams["intitle"] = intitle[0]
	}

	inauthor, exists := queryargs["author"]
	if exists {
		queryparams["inauthor"] = inauthor[0]
	}

	inpublisher, exists := queryargs["publisher"]
	if exists {
		queryparams["inpublisher"] = inpublisher[0]
	}

	bookListModel, err := l.searchConnector.Find(queryparams)
	if err != nil {
		return nil, err
	}

	return bookListModel, nil
}
