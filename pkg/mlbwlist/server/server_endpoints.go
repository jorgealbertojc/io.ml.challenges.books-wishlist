package server

import (
	"github.com/gorilla/mux"
)

func (s *serve) configureServiceEndpoints() error {

	s.configureUsersServiceEndpoints()
	s.configureSigninServiceEndpoint()
	s.configureWishlistsServiceEndpoints()
	s.configureBooksServiceEndpoints()
	s.configureSearchBookServiceEndpoints()

	return nil
}

func (s *serve) printServiceEndpoints() {

	s.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return nil
		}
		methods, err := route.GetMethods()
		if err != nil {
			return nil
		}
		s.logging.Info("Exposing server endpoint %v %s", methods, path)
		return nil
	})
}
