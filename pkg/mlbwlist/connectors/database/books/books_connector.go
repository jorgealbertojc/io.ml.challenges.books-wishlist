package books

import (
	"database/sql"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/commons/logging"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

type Connector interface {
	Insert(book models.Book) error
	Select(userid string, wishlistid string, bookid string) (*models.Book, error)
	SelectByGID(userid string, wishlistid string, gid string) (*models.Book, error)
	List(userid string, wishlist string) (*models.BookList, error)
	Delete(userid string, wishlistid string, bookid string) error
}

type connector struct {
	config  *models.Config
	logging logging.Logging

	db        *sql.DB
	tablename string
}

func New(config *models.Config, logging logging.Logging, database *sql.DB) Connector {

	tablename := "bw_wishlist_books"

	return &connector{
		config:    config,
		logging:   logging,
		db:        database,
		tablename: tablename,
	}
}
