package search

import (
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/commons/logging"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/connectors/google/books/search"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

type Logic interface {
	Find(queryargs map[string][]string) (*models.BookList, error)
}

type logic struct {
	config  *models.Config
	logging logging.Logging

	searchConnector search.Connector
}

func New(config *models.Config, logging logging.Logging, searchConnector search.Connector) Logic {

	return &logic{
		config:          config,
		logging:         logging,
		searchConnector: searchConnector,
	}
}
