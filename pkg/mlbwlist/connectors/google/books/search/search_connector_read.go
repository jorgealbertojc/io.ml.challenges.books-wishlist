package search

import (
	"encoding/json"
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (c *connector) Read(gbookid string) (*models.Book, error) {

	url := fmt.Sprintf("%s/%s", c.config.GSuite.API, gbookid)
	c.logging.Info(url)
	request, _, errs := c.request.Get(url).
		End()
	if errs != nil {
		return nil, errs[0]
	}

	searchResult := SearchResult{}
	err := json.NewDecoder(request.Body).Decode(&searchResult)
	if err != nil {
		return nil, err
	}

	return fromSearchResultToBookItem(searchResult), nil
}
