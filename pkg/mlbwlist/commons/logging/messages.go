package logging

import "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models/config"

type Logging interface {
	Info(message string, args ...interface{}) string
	Error(message string, args ...interface{}) string
}

type logging struct {
	config config.Config
}

func New(config config.Config) Logging {

	return &logging{
		config: config,
	}
}

func (l *logging) Info(template string, args ...interface{}) string {

	return ""
}

func (l *logging) Error(message string, args ...interface{}) string {

	return ""
}
