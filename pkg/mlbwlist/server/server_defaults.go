package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (s *serve) setupDefaultErrorHandlers() {

	s.router.NotFoundHandler = s.setupDefaultNotFoundErrorHandler()
}

func (s *serve) setupDefaultNotFoundErrorHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httperror := models.ServerError{}
		httperror.Status = http.StatusNotFound
		httperror.Spec = models.ServerErrorSpec{
			Type:   "Error",
			Reason: fmt.Sprintf("Route {[%s] %s} was not found in server paths", r.Method, r.URL.Path),
		}
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("Content-Type", "application/json;charset=utf-8")

		s.logging.Error("server has not found route %s", r.URL.Path)
		json.NewEncoder(w).Encode(httperror)
	})
}
