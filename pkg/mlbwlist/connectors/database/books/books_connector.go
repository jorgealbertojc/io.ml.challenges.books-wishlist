package books

import (
	"database/sql"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

type Connector interface {
	Insert(userid string, wishlistid string, book models.Books) error
	Select(userid string, wishlistid string, bookid string) (*models.Books, error)
	List(userid string, wishlistid string) (*models.BooksList, error)
}

type connector struct {
	config *models.Config

	db *sql.DB
}

func New(config *models.Config, database *sql.DB) Connector {

	return &connector{
		config: config,
		db:     database,
	}
}
