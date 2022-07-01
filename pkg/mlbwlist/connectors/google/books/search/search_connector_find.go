package search

import (
	"encoding/json"
	"fmt"
	"strings"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (c *connector) Find(searchTerms map[string]string) (*models.BookList, error) {

	terms := c.buildSearchRequestQueryParams(searchTerms)

	if len(terms) == 0 {
		return nil, fmt.Errorf("search terms cannot be empty")
	}

	url := fmt.Sprintf("%s?%s", c.config.GSuite.API, terms)
	c.logging.Info(url)
	response, _, errs := c.request.Get(url).
		End()
	fmt.Println("response.Request.URL:", response.Request.URL)
	if errs != nil {
		return nil, errs[0]
	}

	if response.StatusCode < 200 && response.StatusCode >= 400 {
		return nil, fmt.Errorf("cannot found any results, please check your connection. Error %d", response.StatusCode)
	}

	list := SearchResultList{Items: []SearchResult{}}
	err := json.NewDecoder(response.Body).Decode(&list)
	if err != nil {
		return nil, err
	}

	bookListModel := models.BookList{Items: []models.Book{}}
	for _, item := range list.Items {
		bookListModel.Items = append(bookListModel.Items, *fromSearchResultToBookItem(item))
	}

	return &bookListModel, nil
}

func (c *connector) buildSearchRequestQueryParams(searchTerms map[string]string) string {

	searchQueryString := ""

	q, exists := searchTerms["q"]
	if exists {
		searchQueryString = fmt.Sprintf("%s%s", searchQueryString, q)
	}

	intitle, exists := searchTerms["intitle"]
	if exists {
		searchQueryString = fmt.Sprintf("%s+intitle:%s", searchQueryString, intitle)
	}

	inauthor, exists := searchTerms["inauthor"]
	if exists {
		searchQueryString = fmt.Sprintf("%s+inauthor:%s", searchQueryString, inauthor)
	}

	inpublisher, exists := searchTerms["inpublisher"]
	if exists {
		searchQueryString = fmt.Sprintf("%s+inpublisher:%s", searchQueryString, inpublisher)
	}

	if searchQueryString[0] == '+' {
		searchQueryString = searchQueryString[1:]
	}

	return fmt.Sprintf("q=%s", strings.ReplaceAll(searchQueryString, " ", "+"))
}
