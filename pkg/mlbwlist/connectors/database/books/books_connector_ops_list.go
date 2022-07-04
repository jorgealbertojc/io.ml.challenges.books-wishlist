package books

import (
	"fmt"

	"io.ml.challenges/io.ml.challenges.books-wishlist/pkg/mlbwlist/models"
)

func (c *connector) List(userid string, wishlistid string) (*models.BookList, error) {

	bookList := models.BookList{Items: []models.Book{}}

	sql := fmt.Sprintf("SELECT _id, meta_user_id, meta_wishlist_id, meta_google_id, spec_title, spec_authors, spec_publisher FROM %s WHERE meta_user_id = '%s' AND meta_wishlist_id = '%s'",
		c.tablename,
		userid, wishlistid)
	c.logging.Info(sql)
	rows, err := c.db.Query(sql)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		bwWishlistBook := BWWishlistBook{}
		rows.Scan(&bwWishlistBook.ID, &bwWishlistBook.MetaUserID, &bwWishlistBook.MetaWishlistID, &bwWishlistBook.MetaGoogleID, &bwWishlistBook.SpecTitle, &bwWishlistBook.SpecAuthors, &bwWishlistBook.SpecPublisher)
		book := fromBWWishlistBookToBookModel(bwWishlistBook)
		bookList.Items = append(bookList.Items, *book)
	}

	return &bookList, nil
}
