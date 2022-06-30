package models

type SearchResultList struct {
	Items []SearchResult `json:"items"`
}

type SearchResult struct {
	ID         string                 `json:"id,omitempty"`
	VolumeInfo SearchResultVolumeInfo `json:"volumeInfo,omitempty"`
}

type SearchResultVolumeInfo struct {
	Title     string   `json:"title,omitempty"`
	Authors   []string `json:"authors,omitempty"`
	Publisher string   `json:"publisher,omitempty"`
}
