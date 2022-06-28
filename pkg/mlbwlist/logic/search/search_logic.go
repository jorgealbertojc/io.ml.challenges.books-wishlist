package search

import (
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/commons/logging"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

type Logic interface {
	Find(arguments map[string][]string) (map[string]string, error)
}

type logic struct {
	config  *models.Config
	logging logging.Logging
}

func New(config *models.Config, logging logging.Logging) Logic {

	return &logic{
		config:  config,
		logging: logging,
	}
}
