package users

import (
	"time"

	"github.com/google/uuid"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (l *logic) Create(user models.UserAccount) (*models.UserAccount, error) {

	user.ID = uuid.New().String()
	user.CreatedAt = int(time.Now().UnixMilli())

	return &user, nil
}
