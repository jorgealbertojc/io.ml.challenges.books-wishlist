package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func (s *serve) validateSigninAuthToken(w http.ResponseWriter, r *http.Request) error {

	token, err := parseBearerTokenFromRequest(r)
	if err != nil {
		return err
	}

	s.logging.Info("BearerToken: %s", token)

	userid := mux.Vars(r)["user"]

	_, err = s.signinConnector.Select(userid, token)
	if err != nil {
		s.logging.Error(err.Error())
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println(r.Response)
		return fmt.Errorf("you are not authorized to execute actions over endpoint {%s}, verify that your token is not expired already", r.URL.Path)
	}

	return nil
}

func parseBearerTokenFromRequest(r *http.Request) (string, error) {

	token := ""
	authorization := r.Header.Get("Authorization")
	auth := strings.Split(authorization, " ")
	if len(auth) != 2 {
		return "", fmt.Errorf("authorization token canno be parsed")
	}

	if len(auth) == 2 {
		token = auth[1]
	}

	return token, nil
}
