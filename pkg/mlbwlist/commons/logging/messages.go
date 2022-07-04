package logging

import (
	"fmt"
	"time"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

type Logging interface {
	Info(template string, args ...interface{}) string
	Error(template string, args ...interface{}) string
}

type logging struct {
	config *models.Config
}

func New(config *models.Config) Logging {

	return &logging{
		config: config,
	}
}

func (l *logging) Info(template string, args ...interface{}) string {

	message := l.println("INFO", fmt.Sprintf(template, args...))
	fmt.Println(message)
	return message
}

func (l *logging) Error(template string, args ...interface{}) string {

	message := l.println("ERROR", fmt.Sprintf(template, args...))
	fmt.Println(message)
	return message
}

func (l *logging) println(messageType string, message string) string {

	now := time.Now().Format(time.RFC3339)
	return fmt.Sprintf("%s - %s -//- [ %s ] :: %s", now, l.config.UUID, messageType, message)
}
