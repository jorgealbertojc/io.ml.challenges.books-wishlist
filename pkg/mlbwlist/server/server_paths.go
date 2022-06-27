package server

const (
	apiversion string = "/v0"

	signinEndpointPath string = "/signin"

	userAccountsEndpointPath string = "/users"
	userAccountEndpointPath  string = "/users/{user}"

	wishListsEndpointPath string = "/users/{user}/wishlists"
	wishListEndpointPath  string = "/users/{user}/wishlists/{wishlist}"

	booksEndpointPath string = "/users/{user}/wishlists/{wishlist}/books"
	bookEndpointPath  string = "/users/{user}/wishlists/{wishlist}/books/{book}"

	searchBookEndpointPath string = "/search"
)
