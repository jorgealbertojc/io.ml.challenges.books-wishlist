package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/mux"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/commons/logging"
	booksconnector "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/connectors/database/books"
	signinconnector "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/connectors/database/signin"
	usersconnector "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/connectors/database/users"
	wishlistconnector "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/connectors/database/wishlist"
	searchconnector "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/connectors/google/books/search"
	bookslogic "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/logic/books"
	searchlogic "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/logic/search"
	signinlogic "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/logic/signin"
	userslogic "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/logic/users"
	wishlistlogic "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/logic/wishlist"
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

	userLogic     userslogic.Logic
	signinLogic   signinlogic.Logic
	wishlistLogic wishlistlogic.Logic
	booksLogic    bookslogic.Logic
	searchLogic   searchlogic.Logic

	userConnector     usersconnector.Connector
	signinConnector   signinconnector.Connector
	wishlistConnector wishlistconnector.Connector
	booksConnector    booksconnector.Connector
	searchConnector   searchconnector.Connector
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

	server := serve{
		config:  config,
		router:  router,
		logging: logging,

		db: database,
	}

	return &server
}

func (s *serve) Configure() error {

	s.setupDefaultErrorHandlers()
	s.configureServiceDBConnectors()
	s.configureServiceLogistics()
	s.configureServiceEndpoints()

	s.printServiceEndpoints()

	return nil
}

func (s *serve) Run() error {

	address := fmt.Sprintf("%s:%d", s.config.Application.AppService.Host, s.config.Application.AppService.Port)
	s.logging.Info("Server start service at address: %s", address)
	return http.ListenAndServe(address, s.router)
}

func (s *serve) configureServiceLogistics() {

	s.userLogic = userslogic.New(s.config, s.logging, s.userConnector)
	s.signinLogic = signinlogic.New(s.config, s.logging, s.signinConnector, s.userConnector)
	s.wishlistLogic = wishlistlogic.New(s.config, s.logging, s.wishlistConnector, s.userConnector)
	s.booksLogic = bookslogic.New(s.config, s.logging, s.booksConnector, s.searchConnector)
	s.searchLogic = searchlogic.New(s.config, s.logging, s.searchConnector)
}

func (s *serve) configureServiceDBConnectors() {

	s.userConnector = usersconnector.New(s.config, s.logging, s.db)
	s.signinConnector = signinconnector.New(s.config, s.logging, s.db)
	s.wishlistConnector = wishlistconnector.New(s.config, s.logging, s.db)
	s.booksConnector = booksconnector.New(s.config, s.logging, s.db)
	s.searchConnector = searchconnector.New(s.config, s.logging)
}
