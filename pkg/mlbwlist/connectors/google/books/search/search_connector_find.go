package search

import (
	"encoding/json"
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (c *connector) Find(searchTerms map[string]string) (*models.SearchResultList, error) {

	url := fmt.Sprintf("%s?q=%s", c.config.GSuite.API, searchTerms["q"])
	c.logging.Info(url)
	response, _, errs := c.request.Get(url).
		End()
	if errs != nil {
		return nil, errs[0]
	}

	if response.StatusCode < 200 && response.StatusCode >= 400 {
		return nil, fmt.Errorf("cannot found any results, please check your connection. Error %d", response.StatusCode)
	}

	list := models.SearchResultList{}
	list.Items = []models.SearchResult{}

	err := json.NewDecoder(response.Body).Decode(&list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
