package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (s *serve) configureBooksServiceEndpoints() {

	createWishlistBookEndpointPath := fmt.Sprintf("%s%s", apiversion, booksEndpointPath)
	s.router.HandleFunc(createWishlistBookEndpointPath, s.manageCreateWishlistBookRequest).
		Methods(http.MethodPost)

	listWishlistBooksEndpointPath := fmt.Sprintf("%s%s", apiversion, booksEndpointPath)
	s.router.HandleFunc(listWishlistBooksEndpointPath, s.manageListWishlistBooksRequest).
		Methods(http.MethodGet)

	readWishlistBookItemEndpointPath := fmt.Sprintf("%s%s", apiversion, bookEndpointPath)
	s.router.HandleFunc(readWishlistBookItemEndpointPath, s.manageReadWishlistBookRequest).
		Methods(http.MethodGet)
}

func (s *serve) manageCreateWishlistBookRequest(w http.ResponseWriter, r *http.Request) {

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

	s.httpJsonResponseManagement(w, model, http.StatusOK)
}

func (s *serve) manageListWishlistBooksRequest(w http.ResponseWriter, r *http.Request) {

	muxvars := mux.Vars(r)
	userid := muxvars["user"]
	wishlistid := muxvars["wishlist"]
	list, err := s.booksLogic.List(userid, wishlistid)
	if err != nil {
		s.httpServiceErrorManagement(w, err.Error())
		return
	}

	s.httpJsonResponseManagement(w, list, http.StatusOK)
}

func (s *serve) manageReadWishlistBookRequest(w http.ResponseWriter, r *http.Request) {

	muxvars := mux.Vars(r)
	userid := muxvars["user"]
	wishlistid := muxvars["wishlist"]
	bookid := muxvars["book"]
	book, err := s.booksLogic.Read(userid, wishlistid, bookid)
	if err != nil {
		s.httpServiceErrorManagement(w, err.Error())
		return
	}

	s.httpJsonResponseManagement(w, book, http.StatusOK)
}
