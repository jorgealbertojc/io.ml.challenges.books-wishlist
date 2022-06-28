package search

func (l *logic) Find(arguments map[string][]string) (map[string]string, error) {

	searchParams := map[string]string{}
	author, exists := arguments["author"]
	if exists {
		searchParams["author"] = author[0]
	}

	title, exists := arguments["title"]
	if exists {
		searchParams["title"] = title[0]
	}

	publisher, exists := arguments["publisher"]
	if exists {
		searchParams["publisher"] = publisher[0]
	}

	return searchParams, nil
}
