package users

import (
	"database/sql"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/commons/logging"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

type Connector interface {
	Insert(userAccountModel models.UserAccount) error
	Select(userid string) (*models.UserAccount, error)
	SelectByUsername(username string) (*models.UserAccount, error)
}

type connector struct {
	config  *models.Config
	logging logging.Logging

	db *sql.DB

	tablename string
}

func New(config *models.Config, logging logging.Logging, database *sql.DB) Connector {

	tablename := "bw_user_accounts"

	return &connector{
		config:    config,
		logging:   logging,
		db:        database,
		tablename: tablename,
	}
}
