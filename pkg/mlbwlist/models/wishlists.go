package models

type Wishlist struct {
	ID   string        `json:"_id,omitempty"`
	Spec *WishlistSpec `json:"spec,omitempty"`
}

type WishlistSpec struct {
	User        string `json:"user,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type WishlistList struct {
	Items []WishlistList `json:"items"`
}
