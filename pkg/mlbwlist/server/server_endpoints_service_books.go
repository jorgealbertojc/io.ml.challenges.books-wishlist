package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (s *serve) configureBooksServiceEndpoints() {

	createBookEndpointPath := fmt.Sprintf("%s%s", apiversion, booksEndpointPath)
	s.router.HandleFunc(createBookEndpointPath, s.manageCreateBookRequest).
		Methods(http.MethodPost)
}

func (s *serve) manageCreateBookRequest(w http.ResponseWriter, r *http.Request) {

	muxvars := mux.Vars(r)
	userid := muxvars["user"]
	wishlistid := muxvars["wishlist"]
	bookModel := models.Books{}
	json.NewDecoder(r.Body).Decode(&bookModel)
	model, err := s.booksLogic.Create(userid, wishlistid, bookModel)
	if err != nil {
		s.httpServiceErrorManagement(w, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json;charset=utf-8")

	json.NewEncoder(w).Encode(model)
}
