package search

import (
	"encoding/json"
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (c *connector) Read(gbookid string) (*models.Book, error) {

	url := fmt.Sprintf("%s/%s", c.config.GSuite.API, gbookid)
	c.logging.Info(url)
	response, _, errs := c.request.Get(url).
		End()
	if errs != nil {
		return nil, errs[0]
	}

	if response.StatusCode < 200 || response.StatusCode >= 400 {
		return nil, fmt.Errorf("cannot found any results, please check the provided book ID. Google Books API Error %d", response.StatusCode)
	}

	searchResult := SearchResult{}
	err := json.NewDecoder(response.Body).Decode(&searchResult)
	if err != nil {
		return nil, err
	}

	return fromSearchResultToBookItem(searchResult), nil
}
