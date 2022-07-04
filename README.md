# io.ml.challenges.books-wishlist

## 1 .- Introducción

Application that exposes a simple API with the following features

### 1.1 .- Users

Manage user accounts creating an account by using `username` and `password` and
read it by using a API Bearer token.

### 1.2 .- Books Wishlists

Manage books wishlist under user account and by using an API Bearer token

### 1.3 .- Books

Manage Books from Books Wishlist, this is just a representation of books items
in Google Books API.

### 1.4 .- Search

This is a tool that finds books in Google Books API with the terms documented in
requirements

## 2 .- Available Endpoints

### 2.1 .- Users

#### 2.1.1 .- Model

```javascript
{
    // _id: the user identifier, is used to find user in database records
    // NOTE: does not need to provide it, the system will assign it automatically
    "_id": "string(match:uuid.v4)",
    // spec: groups the user account specifications
    "spec": {
        // username: the username that the user want to use to identify with the sigin process
        "username": "string",
        // password: the password that the user want to use to be identified with the signin process
        "password": "string"
    }
}
```

#### 2.1.2 .- Endpoints

| Endpoint                  | Method        | Requirements  | Description               |
| :---                      | :---          | :---          | :---                      |
| **`/v0/users`**           | **`POST`**    | _\<none\>_    | Creates user account      |
| **`/v0/users/{user}`**    | **`GET`**     | Bearer token  | Read own user account     |

<a name="2-1-2-user-endpoint-path-params"></a>
**Path Params**

- **`user`**: is the user id (`_id`) assigned by the system in the user account creation

### 2.2 .- Books Wishlist

#### 2.2.1 .- Model

#### 2.2.2 .- Endpoints

### 2.3 .- Books

### 2.4 .- Search

## 3 .- How to Build
