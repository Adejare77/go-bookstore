# BookStore API

The BookStore API is a simple RESTful service for managing book and movie records. It allows users to perform basic CRUD operations without file uploads. Users can add, retrieve, update, and delete information about books or movies, including details like the author, title, and publication.

## Environment Configuration

Create a .env file in the root directory to provide the required environment variables:

```plaintext
USER="yourDatabaseUsername"
MYSQL_PASSWORD="yourDatabasePassword"
DATABASE="bookstore"
PORT="8080"
```

## Setup Instructions

1. Ensure MySQL is installed and running.
2. Create the database named bookstore (or as specified in `.env`) on your MySQL server.
3. Populate the .env file with the necessary values.
4. Build and run the executable.

**Note:** Ensure the database exists on your MySQL server before running the application.

### API Endpoints

- GET /books
    Retrieves all books.

- GET /books/{id}
    Retrieves a specific book by ID.

- POST /books
    Adds a new book entry.

- PUT /books/{id}
    Updates an existing book by ID.

- DELETE /books/{id}
    Deletes a specific book by ID.

### Additional Notes

- This API only stores information on books/movies; it does not handle file uploads.
- Error handling, validation, and proper HTTP status codes are implemented for a consistent API experience.

### Usage Examples

#### Adding a New Book

**Request:**

```http
POST /books
```

**Request Body:**

```json

{
  "title": "Game of Thrones",
  "author": "George R. R. Martin",
  "publication": "1996"
}
```

**Response:**

```json

{
  "id": 1,
  "title": "Game of Thrones",
  "author": "George R. R. Martin",
  "publication": "1996",
}
```

#### Retrieving All Books

**Request:**

```http
GET /books/
```

**Response:**

```python
[
  {
    "id": 1,
    "title": "Game of Thrones",
    "author": "George R. R. Martin",
    "publication": "1996"
  },
  {
    "id": 2,
    "title": "To Kill a Mockingbird",
    "author": "Harper Lee",
    "publication": "1960"
  }
]
```

#### Updating a Book

**Request:**

```http
PUT /books/1
```

**Request Body:**

```json
{
  "title": "A Game of Thrones",
  "author": "George R. R. Martin",
  "publication": "1996"
}
```

**Response:**

```json
{
  "id": 1,
  "title": "A Game of Thrones",
  "author": "George R. R. Martin",
  "publication": "1996",
}
```

#### Deleting a Book

**Request:**

```http
DELETE /book/1
```

**Response:**

```json
{
  "message": "Book deleted successfully"
}
```
