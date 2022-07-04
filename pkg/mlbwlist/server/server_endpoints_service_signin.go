package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (s *serve) configureSigninServiceEndpoint() {

	endpointPath := fmt.Sprintf("%s%s", apiversion, signinEndpointPath)
	method := http.MethodPost
	s.router.HandleFunc(endpointPath, s.manageCreateSigninTokenRequest).
		Methods(method)
}

func (s *serve) manageCreateSigninTokenRequest(w http.ResponseWriter, r *http.Request) {

	userAccountModel := models.UserAccount{}
	json.NewDecoder(r.Body).Decode(&userAccountModel)
	model, err := s.signinLogic.Create(userAccountModel)
	if err != nil {
		s.httpServiceErrorManagement(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.httpJsonResponseManagement(w, model, http.StatusCreated)
}
