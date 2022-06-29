package users

import (
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/commons/logging"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/connectors/database/users"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

type Logic interface {
	Create(user models.UserAccount) (*models.UserAccount, error)
	Read(userid string) (*models.UserAccount, error)
}

type logic struct {
	config  *models.Config
	logging logging.Logging

	db users.Connector
}

func New(config *models.Config, logging logging.Logging, connector users.Connector) Logic {

	return &logic{
		config:  config,
		logging: logging,

		db: connector,
	}
}
