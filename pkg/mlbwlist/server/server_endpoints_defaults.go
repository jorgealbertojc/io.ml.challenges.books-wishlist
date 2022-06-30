package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

func (s *serve) httpJsonResponseManagement(w http.ResponseWriter, model interface{}) {

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(model)
}

func (s *serve) validateSigninAuthToken(w http.ResponseWriter, r *http.Request) error {

	authorization := r.Header.Get("Authorization")
	s.logging.Info("token: %s", authorization)
	auth := strings.Split(authorization, " ")
	token := ""
	if len(auth) == 2 {
		token = auth[1]
	}

	err := s.signinConnector.TokenExists(token)
	if err != nil {
		s.logging.Error(err.Error())
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println(r.Response)
		return fmt.Errorf("you are not authorized to execute actions over endpoint {%s}, verify that your token is not expired already", r.URL.Path)
	}

	return nil
}
