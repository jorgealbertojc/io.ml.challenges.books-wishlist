package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (s *serve) configureWishlistsServiceEndpoints() {

	endpointPath := fmt.Sprintf("%s%s", apiversion, wishListsEndpointPath)
	method := http.MethodPost
	s.router.HandleFunc(endpointPath, s.manageCreateWishlistsRequest).
		Methods(method)
}

func (s *serve) manageCreateWishlistsRequest(w http.ResponseWriter, r *http.Request) {

	userid := mux.Vars(r)["user"]
	wishlistModel := models.Wishlist{}
	json.NewDecoder(r.Body).Decode(&wishlistModel)
	if wishlistModel.Spec == nil {
		wishlistModel.Spec = &models.WishlistSpec{}
	}
	wishlistModel.Spec.User = userid
	model, err := s.wishlistLogic.Create(wishlistModel)
	if err != nil {
		s.httpServiceErrorManagement(w, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json;charset=utf-8")

	json.NewEncoder(w).Encode(model)
}
