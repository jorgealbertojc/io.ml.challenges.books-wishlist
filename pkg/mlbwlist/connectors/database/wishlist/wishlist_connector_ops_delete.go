package wishlist

import "fmt"

func (c *connector) Delete(userid string, wishlistid string) error {

	sql := fmt.Sprintf("DELETE FROM %s WHERE meta_user_id = '%s' AND _id = '%s'",
		c.tablename,
		userid, wishlistid)
	c.logging.Info(sql)
	_, err := c.db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
