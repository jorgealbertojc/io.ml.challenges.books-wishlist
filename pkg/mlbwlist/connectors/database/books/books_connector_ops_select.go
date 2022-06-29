package books

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (c *connector) Select(userid string, wishlistid string, bookid string) (*models.Book, error) {

	bwWishlistBook := BWWishlistBook{}
	sql := fmt.Sprintf("SELECT _id, meta_user_id, meta_wishlist_id, meta_google_id, spec_title, spec_authors, spec_publisher FROM %s WHERE meta_user_id = '%s' AND meta_wishlist_id = '%s' AND _id = '%s'",
		c.tablename,
		userid, wishlistid, bookid)
	c.logging.Info(sql)
	err := c.db.QueryRow(sql).
		Scan(&bwWishlistBook.ID, &bwWishlistBook.MetaUserID, &bwWishlistBook.MetaWishlistID, &bwWishlistBook.MetaGoogleID, &bwWishlistBook.SpecTitle, &bwWishlistBook.SpecAuthors, &bwWishlistBook.SpecPublisher)
	if err != nil {
		return nil, err
	}

	return fromBWWishlistBookToBookModel(bwWishlistBook), nil
}
