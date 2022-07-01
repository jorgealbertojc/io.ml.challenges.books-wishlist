package server

import (
	"fmt"
	"net/http"
)

func (s *serve) configureSearchBookServiceEndpoints() {

	executeSearchBookEndpointPath := fmt.Sprintf("%s%s", apiversion, searchBookEndpointPath)
	s.router.HandleFunc(executeSearchBookEndpointPath, s.manageExecuteSearchBookEndpointRequest).
		Methods(http.MethodGet)
}

func (s *serve) manageExecuteSearchBookEndpointRequest(w http.ResponseWriter, r *http.Request) {

	token, err := parseBearerTokenFromRequest(r)
	if err != nil {
		s.httpServiceErrorManagement(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if len(token) == 0 {
		s.httpServiceErrorManagement(w, "you are not allowed to perform search operation without authentication", http.StatusUnauthorized)
		return
	}

	arguments := r.URL.Query()
	result, err := s.searchLogic.Find(arguments)
	if err != nil {
		s.httpServiceErrorManagement(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.httpJsonResponseManagement(w, result, http.StatusOK)
}
