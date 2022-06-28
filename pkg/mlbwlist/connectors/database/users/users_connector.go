package users

import (
	"database/sql"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

type Connector interface {
	Insert(userAccountModel models.UserAccount) error
	Select(userid string) (*models.UserAccount, error)
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
