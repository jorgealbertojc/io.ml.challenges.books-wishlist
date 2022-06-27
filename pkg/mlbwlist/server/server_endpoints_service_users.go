package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (s *serve) configureUsersServiceEndpoints() {

	endpointPath := fmt.Sprintf("%s%s", apiversion, userAccountsEndpointPath)
	method := http.MethodPost
	s.router.HandleFunc(endpointPath, s.manageCreateUserAccount).
		Methods(method)
}

func (s *serve) manageCreateUserAccount(w http.ResponseWriter, r *http.Request) {

	userAccountModel := models.UserAccount{}
	json.NewDecoder(r.Body).Decode(&userAccountModel)
	model, err := s.userLogic.Create(userAccountModel)
	if err != nil {
		s.httpServiceErrorManagement(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(model)
}
