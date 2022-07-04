package books

import "fmt"

func (c *connector) Delete(userid string, wishlistid string, bookid string) error {

	sql := fmt.Sprintf("DELETE FROM %s WHERE meta_user_id = '%s' AND meta_wishlist_id = '%s' AND _id = '%s'",
		c.tablename,
		userid, wishlistid, bookid)
	c.logging.Info(sql)
	_, err := c.db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
