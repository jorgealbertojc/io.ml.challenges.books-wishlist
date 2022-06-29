package wishlist

import (
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/commons/logging"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/connectors/database/wishlist"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

type Logic interface {
	Create(wishlistModel models.Wishlist) (*models.Wishlist, error)
	Read(userid string, wishlistid string) (*models.Wishlist, error)
	List(userid string) (*models.WishlistList, error)
}

type logic struct {
	config  *models.Config
	logging logging.Logging

	db wishlist.Connector
}

func New(config *models.Config, logging logging.Logging, db wishlist.Connector) Logic {

	return &logic{
		config:  config,
		logging: logging,
		db:      db,
	}
}
