package users

import "io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"

func (l *logic) Read(userid string) (*models.UserAccount, error) {

	l.logging.Info("requesting user id {%s}", userid)
	user, err := l.db.Select(userid)
	if err != nil {
		return nil, err
	}

	return user, nil
}
