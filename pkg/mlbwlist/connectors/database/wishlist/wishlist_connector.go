package wishlist

import (
	"database/sql"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

type Connector interface {
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
