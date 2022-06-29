package signin

import (
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/commons/logging"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/connectors/database/signin"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/connectors/database/users"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

type Logic interface {
	Create(userAccountModel models.UserAccount) (*models.Signin, error)
}

type logic struct {
	config  *models.Config
	logging logging.Logging

	db     signin.Connector
	userdb users.Connector
}

func New(config *models.Config, logging logging.Logging, db signin.Connector, userdb users.Connector) Logic {

	return &logic{
		config:  config,
		logging: logging,

		db:     db,
		userdb: userdb,
	}
}
