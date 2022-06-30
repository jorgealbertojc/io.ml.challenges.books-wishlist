package models

type Book struct {
	ID   string    `json:"_id,omitempty"`
	Meta *BookMeta `json:"meta,omitempty"`
	Spec *BookSpec `json:"spec,omitempty"`
}

type BookMeta struct {
	UserID     string `json:"user,omitempty"`
	WishlistID string `json:"wishlist,omitempty"`
	GoogleID   string `json:"gid,omitempty"`
}

type BookSpec struct {
	Title     string   `json:"title,omitempty"`
	Authors   []string `json:"authors,omitempty"`
	Publisher string   `json:"publisher,omitempty"`
}

type BookList struct {
	Items []Book `json:"items,omitempty"`
}
