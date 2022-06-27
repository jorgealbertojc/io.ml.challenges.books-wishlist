package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/mux"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/commons/logging"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/logic/books"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/logic/search"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/logic/users"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/logic/wishlist"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

type Server interface {
	Configure() error
	Run() error
}

type serve struct {
	config  *models.Config
	router  *mux.Router
	logging logging.Logging

	db *sql.DB

	userLogic     users.Logic
	wishlistLogic wishlist.Logic
	booksLogic    books.Logic
	searchLogic   search.Logic
}

func New(config *models.Config) Server {

	router := mux.NewRouter().StrictSlash(true)
	logging := logging.New(config)

	database, err := sql.Open(config.Database.Type, config.Database.Filepath)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("ERROR: cannot initialize database connection")
		os.Exit(4)
	}

	userLogic := users.New(config, logging, database)

	server := serve{
		config:  config,
		router:  router,
		logging: logging,

		db: database,

		userLogic: userLogic,
	}

	return &server
}

func (s *serve) Configure() error {

	s.setupUsersEndpoints()
	s.setupDefaultErrorHandlers()

	return nil
}

func (s *serve) Run() error {

	address := fmt.Sprintf("%s:%d", s.config.Application.AppService.Host, s.config.Application.AppService.Port)
	s.logging.Info("server has starting serve at address: %s", address)
	http.Handle("/", s.router)
	return http.ListenAndServe(address, s.router)
}

func (s *serve) httpServiceErrorManagement(w http.ResponseWriter, message string) {
	errors := models.ServerError{}
	errors.Status = http.StatusBadRequest
	errors.Spec = models.ServerErrorSpec{
		Type:   "Error",
		Reason: message,
	}
	s.logging.Error(message)
	json.NewEncoder(w).Encode(errors)
}
