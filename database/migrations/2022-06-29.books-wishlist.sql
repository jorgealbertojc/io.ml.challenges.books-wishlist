
DROP TABLE IF EXISTS bw_user_accounts;
CREATE TABLE IF NOT EXISTS bw_user_accounts (
    _id VARCHAR(36),
    spec_username VARCHAR(100),
    spec_password VARCHAR(32)
);

DROP TABLE IF EXISTS bw_user_signin_tokens;
CREATE TABLE IF NOT EXISTS bw_user_signin_tokens (
    _id VARCHAR(36),
    meta_user_id VARCHAR(36),
    spec_token_hash TEXT
);

DROP TABLE IF EXISTS bw_wishlists;
CREATE TABLE IF NOT EXISTS bw_wishlists (
    _id VARCHAR(36),
    meta_user_id VARCHAR(36),
    spec_name VARCHAR(300),
    spec_description TEXT
);

DROP TABLE IF EXISTS bw_wishlist_books;
CREATE TABLE IF NOT EXISTS bw_wishlist_books (
    _id VARCHAR(36),
    meta_user_id VARCHAR(36),
    meta_wishlist_id VARCHAR(36),
    meta_google_id TEXT,
    spec_title VARCHAR(300),
    spec_authors TEXT,
    spec_publisher VARCHAr(300)
);
