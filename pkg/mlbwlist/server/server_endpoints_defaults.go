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
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusNotFound)

		s.logging.Error(httperror.Spec.Reason)
		json.NewEncoder(w).Encode(httperror)
	})
}

func (s *serve) httpServiceErrorManagement(w http.ResponseWriter, message string, httpstatus int) {
	errors := models.ServerError{}
	errors.Status = httpstatus
	errors.Spec = models.ServerErrorSpec{
		Type:   "Error",
		Reason: message,
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(errors.Status)

	s.logging.Error(errors.Spec.Reason)
	json.NewEncoder(w).Encode(errors)
}

func (s *serve) httpJsonResponseManagement(w http.ResponseWriter, model interface{}, httpstatus int) {

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(httpstatus)

	if model != nil {
		json.NewEncoder(w).Encode(model)
	}
}
