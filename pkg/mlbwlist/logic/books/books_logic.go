package books

import (
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/commons/logging"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/connectors/database/books"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/connectors/google/books/search"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

type Logic interface {
	Create(bookModel models.Book) (*models.Book, error)
	List(userid string, wishlistid string) (*models.BookList, error)
	Read(userid string, wishlistid string, bookid string) (*models.Book, error)
}

type logic struct {
	config  *models.Config
	logging logging.Logging

	db              books.Connector
	searchConnector search.Connector
}

func New(config *models.Config, logging logging.Logging, db books.Connector, searchConnector search.Connector) Logic {

	return &logic{
		config:          config,
		logging:         logging,
		db:              db,
		searchConnector: searchConnector,
	}
}
