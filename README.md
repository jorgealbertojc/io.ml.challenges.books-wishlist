# io.ml.challenges.books-wishlist

## 0 .- Requirements

- `go 1.18`
- `docker 20.10.17`
- `bash 4.4.20(1)-release`
- `gnu-make 4.1`
- `git 2.17.1`
- `sqlite 3.22.0`

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
        "username": "string(maxlength:100)",
        // password: the password that the user want to use to be identified with the signin process
        "password": "string"
    }
}
```

#### 2.1.2 .- Endpoints

| Endpoint                  | Method        | Requirements          | Description               |
| :---                      | :---          | :---                  | :---                      |
| **`/v0/users`**           | **`POST`**    | _\<none\>_            | Creates user account      |
| **`/v0/users/{user}`**    | **`GET`**     | Bearer token header   | Read own user account     |

<a name="2-1-2-user-endpoint-path-params"></a>
**Path Params**

- **`user`**: is the user id (`_id`) assigned by the system in the user account creation

**Example Request**

```bash
curl --location --silent --show-error \
    --header "Authorization: Bearer ${API_TOKEN}" \
    --request GET \
    --url ${API_HOST}/v0/users/${user}
```

### 2.2 .- Signin

#### 2.2.1 .- Model

**Input**

```javascript
{
    // spec: groups the user account authentication requirements
    "spec": {
        // username: the username the user used on user account creation
        "username": "string(maxlength:100)",
        // password: the password the user used on user account creation
        "password": "string"
    }
}
```

**Output**

```javascript
{
    // _id: unique identifier for token in database records
    "_id": "string(match:uuid.v4)",
    // meta: groups the metadata information about the entity
    "meta": {
        // user: is the user unique identifier assignated on user account creation
        "user": "string(match:uuid.v4)"
    },
    // spec: groups the entity specifications
    "spec": {
        // token: bearer token required in protected endpoints
        "token": "string(length:120)"
    }
}
```

#### 2.2.2 .- Endpoints

| Endpoint          | Method        | Requirements  | Description                   |
| :---              | :---          | :---          | :---                          |
| **`/v0/signin`**  | **`POST`**    | _\<none\>_    | Creates user bearer token     |

**Example Request**

```bash
curl --location --silent --show-error \
    --request POST \
    --url ${API_HOST}/v0/signin \
    --data-raw '{
        "spec": {
            "username": "${API_USERNAME}",
            "password": "${API_PASSWORD}"
        }
    }'
```

### 2.3 .- Books Wishlist

#### 2.3.1 .- Model

```javascript
{
    // _id: wishlist unique identifier asigned automatically by the system
    "_id": "string(match:uuid.v4)",
    // meta: groups the metadata information about the entity
    "meta": {
        // user: is the user unique identifier assignated on user account creation
        "user": "string(match:uuid.v4)"
    },
    // spec: groups the entity specifications
    "spec": {
        // name: a friendly name to identify wishlist
        "name": "string(maxlength:100)",
        // description: a short description to describe the wishlist target
        "description": "string"
    }
}
```

#### 2.3.2 .- Endpoints

| Endpoint                                      | Method        | Requirements          | Description                                               |
| :---                                          | :---          | :---                  | :---                                                      |
| **`/v0/users/{user}/wishlists`**              | **`POST`**    | Bearer token header   | Creates user books wishlist                               |
| **`/v0/users/{user}/wishlists`**              | **`GET`**     | Bearer token header   | List user books wishlists                                 |
| **`/v0/users/{user}/wishlists/{wishlist}`**   | **`GET`**     | Bearer token header   | Describe user wishlist identified with user and list id   |
| **`/v0/users/{user}/wishlists/{wishlist}`**   | **`DELETE`**  | Bearer token header   | Deletes an already existent book wishlist                 |

**Path Params**

- `user`: user unique identifier
- `wishlist`: wishlist unique identifier (_this is automaticaly generated_)

**Example Request**

```bash
curl --location --silent --show-error \
    --request POST \
    --url ${API_HOST}/v0/users/${user}/wishlists \
    --data-raw '{
        "spec": {
            "name": "My Books Wishlist",
            "description": "A books wishlist that I love to read"
        }
    }'
```

### 2.4 .- Books

#### 2.4.1 .- Model

**Input**

```javascript
{
    // meta: groups the metadata information about the entity
    "meta": {
        // gid: Google Book ID that references the information about the book
        "gid": "string(match:googlebookid)"
    }
}
```

> _NOTE: With `meta.gid` the API will request book information from Google Books API_

**Output**

```javascript
{
    // _id: book unique identifier asigned automatically by the system
    "_id": "string(match:uuid.v4)",
    // meta: groups the metadata information about the entity
    "meta": {
        // user: user unique identifier
        "user": "string(match:uuid.v4)",
        // wishlist: wishlist unique identifier
        "wishlist": "string(match:uuid.v4)",
        // gid: Google Book Identifier
        "gid": "string(match:googlebookid)"
    },
    // spec: groups the entity specifications
    "spec": {
        // title: book title
        "title": "string",
        // authors: book authors list
        "authors": [
            "string"
        ],
        // publisher: book publisher
        "publisher": "string"
    }
}
```

#### 2.4.2 .- Endpoints

| Endpoint                                                  | Method        | Requirements          | Description               |
| :---                                                      | :---          | :---                  | :---                      |
| **`/v0/users/{user}/wishlists/{wishlist}/books`**         | **`POST`**    | Bearer token header   | Creates wishlist book     |
| **`/v0/users/{user}/wishlists/{wishlist}/books`**         | **`GET`**     | Bearer token header   | List wishlist books     |
| **`/v0/users/{user}/wishlists/{wishlist}/books/{book}`**  | **`GET`**     | Bearer token header   | Describes wishlist book     |
| **`/v0/users/{user}/wishlists/{wishlist}/books/{book}`**  | **`DELETE`**  | Bearer token header   | Delete wishlist book     |

**Path Params**

- `user`: user unique identifier
- `wishlist`: wishlist unique identifier
- `book`: book unique identifier (_this is automatically generated_)

**Example Request**

```bash
curl --location --silent --show-error \
    --request POST \
    --url ${API_HOST}/v0/users/${user}/wishlists/${wishlist}/books \
    --data-raw '{
        "meta": {
            "gid": "6klRAAAACAAJ"
        }
    }'
```

### 2.5 .- Search

#### 2.5.1 .- Model

**Output**

```javascript
{
    // items: list of books results from the search request
    "items": [
        {
            // _id: book unique identifier asigned automatically by the system
            "_id": "string(match:uuid.v4)",
            // meta: groups the metadata information about the entity
            "meta": {
                // gid: Google Book Identifier
                "gid": "string(match:googlebookid)"
            },
            // spec: groups the entity specifications
            "spec": {
                // title: book title
                "title": "string",
                // authors: book authors list
                "authors": [
                    "string"
                ],
                // publisher: book publisher
                "publisher": "string"
            }
        }
    ]
}
```

#### 2.5.2 .- Endpoints

| Endpoint          | Method    | Requirements          | Description                            |
| :---              | :---      | :---                  | :---                                   |
| **`/v0/search`**  | **`GET`** | Bearer token header   | Search a book in the Google Books API  |

**Query String Params**

| Param             | Type          | Description                                                                                   |
| :---              | :---          | :---                                                                                          |
| **`q`**           | **`string`**  | A general search instruction, this can contains title, author, publisher mixed in the request |
| **`title`**       | **`string`**  | Filter to find books only from specificed title                                               |
| **`author`**      | **`string`**  | Filter to find books only from specificed author                                              |
| **`publisher`**   | **`string`**  | Filter to find books only from specificed publisher                                           |

**Example Request**

```bash
curl --location --silent --show-error \
    --request GET \
    --url ${API_HOST}/v0/search?q=la+caza+de+nimrod&author=Charles+Sheffield&publisher=Editorial+CLIE
```

## 3 .- How to Build

### 3.1 .- Clonning

The repository have an specific way to build and with this in mind it requires a specific route to be cloned.

1. In your `${GOPATH}` directory create the following directory
    ```bash
    mkdir -p ${GOPATH}/src/io.ml.challenges
    ```
2. Move to created directory
    ```bash
    cd ${GOPATH}/src/io.ml.challenges
    ```
3. Clone repository inside the directory created in the directory created in step above
    ```bash
    git clone git@github.com:jorgealbertojc/io.ml.challenges.books-wishlist.git
    ```

### 3.2 .- Building

#### 3.2.1 .- By Hand with Dockerfile

To build application by hand only execute the following commands

```bash
cd ${GOPATH}/src/io.ml.challenges/io.ml.challenges.books-wishlist && \
docker build --tag ${YOUR_REPOSITORY_NAME}:${YOUR_TAG_ID} --file $(pwd)/Dockerfile .
```

#### 3.2.2 .- With Makefile

The Makefile is prepared to build the application and store it inside docker image jus by executing the following command:

```bash
cd ${GOPATH}/src/io.ml.challenges/io.ml.challenges.books-wishlist \
&& make docker-build
```

> _After the make recipe finished a docker image is built and ready to be launched **`io.ml.challenges.books-wishlist:v1.0.0-rc`**_

## 4 .- How to Run

### 4.1 .- Configuration

First thing is to create a configuration file, this configuration file requires
to contains the information to manage the information that will be interpreted
by the daemon compiled in the building section "_Building_".

1. Creates a file with name `config.yml`, a directory can be used to separate configuration from another things like `configurations/config.yml`:
    ```bash
    mkdir -p configurations && \
    touch configurations/config.yml
    ```
2. Add the following contents to `configurations/config.yml`:
    ```yml
    ---
    # app: groups the app configuration
    app:
        # configversion: the configuration file version
        configversion: "1.0"
        # service: groups the information about service exposed port and host
        service:
            # port: is the port where the service will be exposed
            port: 80
            # host: the host where the service will be exposed
            host: "0.0.0.0"

    # database: groups the database configuration
    database:
        # type: represents the type of database (by this moment just sqlite3 is supported)
        type: "sqlite3"
        # filepath: for sqlite3 this is the default path to the database file
        filepath: "/var/local/mlbwlistd/database.db"

    # gsuite: groups the google suite configuration
    gsuite:
        api: "https://www.googleapis.com/books/v1/volumes"
    ```

> _NOTE: the Google Books API does not requires a token nor authentication_

### 4.2 .- Execution

The project have a default configuration filepath after docker image is built, the configuration requires to be stored in the docker image in the following path:

```
/etc/mlbwlistd/configs/config.yml
```

So, to run the docker image with the correct arguments needs to mount a volume where the config.yml is stored, just run the following command:

```bash
docker run \
    -p 8084:80 \
    -v /path/to/directory/where/configyml/is/stored:/etc/mlbwlistd/configs:ro \
    io.ml.challenges.books-wishlist:v1.0.0-rc
```

After run you'll se something like:

```
[ INFO ] :: Exposing server endpoint [POST] /v0/users
[ INFO ] :: Exposing server endpoint [GET] /v0/users/{user}
[ INFO ] :: Exposing server endpoint [POST] /v0/signin
[ INFO ] :: Exposing server endpoint [POST] /v0/users/{user}/wishlists
[ INFO ] :: Exposing server endpoint [GET] /v0/users/{user}/wishlists/{wishlist}
[ INFO ] :: Exposing server endpoint [GET] /v0/users/{user}/wishlists
[ INFO ] :: Exposing server endpoint [DELETE] /v0/users/{user}/wishlists/{wishlist}
[ INFO ] :: Exposing server endpoint [POST] /v0/users/{user}/wishlists/{wishlist}/books
[ INFO ] :: Exposing server endpoint [GET] /v0/users/{user}/wishlists/{wishlist}/books
[ INFO ] :: Exposing server endpoint [GET] /v0/users/{user}/wishlists/{wishlist}/books/{book}
[ INFO ] :: Exposing server endpoint [DELETE] /v0/users/{user}/wishlists/{wishlist}/books/{book}
[ INFO ] :: Exposing server endpoint [GET] /v0/search
```

finally make a request over any of the documented API endpoints to see how it works !
