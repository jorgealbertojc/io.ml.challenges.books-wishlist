package models

type Books struct {
	ID        string `json:"_id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Published string `json:"publisher"`
}

type BooksList struct {
	Items []Books `json:"items"`
}
