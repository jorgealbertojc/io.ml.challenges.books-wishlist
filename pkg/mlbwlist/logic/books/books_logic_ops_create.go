package books

import (
	"fmt"

	"github.com/google/uuid"
	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (l *logic) Create(bookModel models.Book) (*models.Book, error) {

	err := l.validateEntityParents(bookModel)
	if err != nil {
		return nil, err
	}

	storedBookModel, _ := l.validateBookExistenceByGID(bookModel)
	if storedBookModel != nil {
		return nil, fmt.Errorf("book identified with {%s / %s / gid: %s} already exists in our records", storedBookModel.Meta.UserID, storedBookModel.Meta.WishlistID, storedBookModel.Meta.GoogleID)
	}

	bookModel.ID = uuid.New().String()

	searchBookModel, err := l.searchapi.Read(bookModel.Meta.GoogleID)
	if err != nil {
		return nil, err
	}

	bookModel.Spec.Title = searchBookModel.Spec.Title
	bookModel.Spec.Authors = searchBookModel.Spec.Authors
	bookModel.Spec.Publisher = searchBookModel.Spec.Publisher

	err = l.db.Insert(bookModel)
	if err != nil {
		return nil, err
	}

	return &bookModel, nil
}
