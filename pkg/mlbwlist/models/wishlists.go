package models

type Wishlist struct {
	ID   string        `json:"_id,omitempty"`
	Meta *WishlistMeta `json:"meta,omitempty"`
	Spec *WishlistSpec `json:"spec,omitempty"`
}

type WishlistMeta struct {
	UserID string `json:"user,omitempty"`
}

type WishlistSpec struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type WishlistList struct {
	Items []Wishlist `json:"items"`
}
