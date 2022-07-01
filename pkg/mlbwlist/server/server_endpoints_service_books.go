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

	deleteWishlistBookItemEndpointPath := fmt.Sprintf("%s%s", apiversion, bookEndpointPath)
	s.router.HandleFunc(deleteWishlistBookItemEndpointPath, s.manageDeleteWishlistBookRequest).
		Methods(http.MethodDelete)
}

func (s *serve) manageCreateWishlistBookRequest(w http.ResponseWriter, r *http.Request) {

	err := s.validateSigninAuthToken(w, r)
	if err != nil {
		s.httpServiceErrorManagement(w, err.Error(), http.StatusUnauthorized)
		return
	}

	muxvars := mux.Vars(r)
	userid := muxvars["user"]
	wishlistid := muxvars["wishlist"]
	bookModel := models.Book{}
	json.NewDecoder(r.Body).Decode(&bookModel)
	if bookModel.Meta == nil {
		bookModel.Meta = &models.BookMeta{}
	}
	if bookModel.Spec == nil {
		bookModel.Spec = &models.BookSpec{}
	}
	bookModel.Meta.UserID = userid
	bookModel.Meta.WishlistID = wishlistid
	model, err := s.booksLogic.Create(bookModel)
	if err != nil {
		s.httpServiceErrorManagement(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.httpJsonResponseManagement(w, model, http.StatusCreated)
}

func (s *serve) manageListWishlistBooksRequest(w http.ResponseWriter, r *http.Request) {

	err := s.validateSigninAuthToken(w, r)
	if err != nil {
		s.httpServiceErrorManagement(w, err.Error(), http.StatusUnauthorized)
		return
	}

	muxvars := mux.Vars(r)
	userid := muxvars["user"]
	wishlistid := muxvars["wishlist"]
	list, err := s.booksLogic.List(userid, wishlistid)
	if err != nil {
		s.httpServiceErrorManagement(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.httpJsonResponseManagement(w, list, http.StatusOK)
}

func (s *serve) manageReadWishlistBookRequest(w http.ResponseWriter, r *http.Request) {

	err := s.validateSigninAuthToken(w, r)
	if err != nil {
		s.httpServiceErrorManagement(w, err.Error(), http.StatusUnauthorized)
		return
	}

	muxvars := mux.Vars(r)
	userid := muxvars["user"]
	wishlistid := muxvars["wishlist"]
	bookid := muxvars["book"]
	book, err := s.booksLogic.Read(userid, wishlistid, bookid)
	if err != nil {
		s.httpServiceErrorManagement(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.httpJsonResponseManagement(w, book, http.StatusOK)
}

func (s *serve) manageDeleteWishlistBookRequest(w http.ResponseWriter, r *http.Request) {

	err := s.validateSigninAuthToken(w, r)
	if err != nil {
		s.httpServiceErrorManagement(w, err.Error(), http.StatusUnauthorized)
		return
	}

	muxvars := mux.Vars(r)
	userid := muxvars["user"]
	wishlistid := muxvars["wishlist"]
	bookid := muxvars["book"]
	err = s.booksLogic.Delete(userid, wishlistid, bookid)
	if err != nil {
		s.httpServiceErrorManagement(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.httpJsonResponseManagement(w, nil, http.StatusNoContent)
}
