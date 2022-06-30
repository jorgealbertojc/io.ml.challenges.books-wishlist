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

	arguments := r.URL.Query()
	result, err := s.searchLogic.Find(arguments)
	if err != nil {
		s.httpServiceErrorManagement(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.httpJsonResponseManagement(w, result)
}
