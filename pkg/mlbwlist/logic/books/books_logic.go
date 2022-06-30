package books

import (
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/commons/logging"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/connectors/database/books"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/connectors/database/users"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/connectors/database/wishlist"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/connectors/google/books/search"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

type Logic interface {
	Create(bookModel models.Book) (*models.Book, error)
	List(userid string, wishlistid string) (*models.BookList, error)
	Read(userid string, wishlistid string, bookid string) (*models.Book, error)
	Delete(userid string, wishlistid string, bookid string) error
}

type logic struct {
	config  *models.Config
	logging logging.Logging

	db         books.Connector
	searchapi  search.Connector
	userdb     users.Connector
	wishlistdb wishlist.Connector
}

func New(config *models.Config, logging logging.Logging, db books.Connector, searchapi search.Connector, userdb users.Connector, wishlistdb wishlist.Connector) Logic {

	return &logic{
		config:     config,
		logging:    logging,
		db:         db,
		searchapi:  searchapi,
		userdb:     userdb,
		wishlistdb: wishlistdb,
	}
}
