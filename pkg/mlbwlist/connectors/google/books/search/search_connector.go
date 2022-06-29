package search

import (
	"github.com/parnurzeal/gorequest"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/commons/logging"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

type Connector interface {
	Find(searchTerms map[string]string) (*models.SearchResultList, error)
	Read(gbookid string) (*models.SearchResult, error)
}

type connector struct {
	config  *models.Config
	logging logging.Logging

	request *gorequest.SuperAgent
}

func New(config *models.Config, logging logging.Logging) Connector {

	request := gorequest.New()

	return &connector{
		config:  config,
		logging: logging,

		request: request,
	}
}
