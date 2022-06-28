package wishlist

import (
	"database/sql"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/commons/logging"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

type Logic interface {
	Create(wishlistModel models.Wishlist) (*models.Wishlist, error)
	List(userid string) ([]models.Wishlist, error)
}

type logic struct {
	config  *models.Config
	logging logging.Logging

	db *sql.DB
}

func New(config *models.Config, logging logging.Logging, db *sql.DB) Logic {

	return &logic{
		config:  config,
		logging: logging,
		db:      db,
	}
}
