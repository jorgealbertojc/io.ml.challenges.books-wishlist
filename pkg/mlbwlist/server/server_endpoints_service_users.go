package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (s *serve) configureUsersServiceEndpoints() {

	createUserAccountEndpointPath := fmt.Sprintf("%s%s", apiversion, userAccountsEndpointPath)
	s.router.HandleFunc(createUserAccountEndpointPath, s.manageCreateUserAccountRequest).
		Methods(http.MethodPost)

	readUserAccountEndpointPath := fmt.Sprintf("%s%s", apiversion, userAccountEndpointPath)
	s.router.HandleFunc(readUserAccountEndpointPath, s.manageReadUserAccountRequest).
		Methods(http.MethodGet)
}

func (s *serve) manageCreateUserAccountRequest(w http.ResponseWriter, r *http.Request) {

	userAccountModel := models.UserAccount{}
	json.NewDecoder(r.Body).Decode(&userAccountModel)
	model, err := s.userLogic.Create(userAccountModel)
	if err != nil {
		s.httpServiceErrorManagement(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.httpJsonResponseManagement(w, model)
}

func (s *serve) manageReadUserAccountRequest(w http.ResponseWriter, r *http.Request) {

	err := s.validateSigninAuthToken(w, r)
	if err != nil {
		s.httpServiceErrorManagement(w, err.Error(), http.StatusUnauthorized)
		return
	}

	userid := mux.Vars(r)["user"]
	model, err := s.userLogic.Read(userid)
	if err != nil {
		s.httpServiceErrorManagement(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.httpJsonResponseManagement(w, model)
}
