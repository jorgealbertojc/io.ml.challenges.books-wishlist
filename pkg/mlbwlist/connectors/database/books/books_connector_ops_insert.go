package books

import (
	"encoding/json"
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (c *connector) Insert(book models.Book) error {

	authors, _ := json.Marshal(book.Spec.Authors)
	sql := fmt.Sprintf("INSERT INTO %s(_id, meta_user_id, meta_wishlist_id, meta_google_id, spec_title, spec_authors, spec_publisher) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s')",
		c.tablename,
		book.ID,
		book.Meta.UserID, book.Meta.WishlistID, book.Meta.GoogleID,
		book.Spec.Title, authors, book.Spec.Publisher)
	c.logging.Info(sql)
	_, err := c.db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
