package search

import (
	"github.com/parnurzeal/gorequest"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

type Connector interface {
	Find(searchTerms map[string]string) error
}

type connector struct {
	config *models.Config

	request *gorequest.SuperAgent
}

func New(config *models.Config) Connector {

	request := gorequest.New()

	return &connector{
		config: config,

		request: request,
	}
}
