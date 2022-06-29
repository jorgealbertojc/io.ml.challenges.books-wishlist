package signin

import (
	"database/sql"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/commons/logging"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

type Connector interface {
	Insert(signinModel models.Signin) error
	Select(userid string, token string) (*models.Signin, error)
}

type connector struct {
	config  *models.Config
	logging logging.Logging

	db *sql.DB
}

func New(config *models.Config, logging logging.Logging, database *sql.DB) Connector {

	return &connector{
		config:  config,
		db:      database,
		logging: logging,
	}
}
