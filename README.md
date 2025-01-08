# Simple RESTful API Server

This project is a simple RESTful API built using Fiber. It provides basic CRUD (Create, Read, Update, Delete) functionality for managing items in an in-memory database. The API is thread-safe, utilizing a lock to handle concurrent access.

---

## Features

- **Get All Items**: Retrieve a list of all stored items.
- **Get Item by ID**: Fetch a specific item by its ID.
- **Create Item**: Add a new item to the database.
- **Update Item**: Modify an existing item using its ID.
- **Delete Item**: Remove an item using its ID.

---

## Endpoints

### 1. Get All Items

- **URL**: `/api/v1/items`
- **Method**: `GET`
- **Response**: List of all items.

### 2. Get Item by ID

- **URL**: `/api/v1/item/<id>`
- **Method**: `GET`
- **Response**:
  `200 OK with the item if found.`\
  `404 Not Found if the item does not exist.`\

### 3. Create an Item

- **URL**: `/api/v1/item`
- **Method**: `POST`
- **Response**:
  JSON object with item details.

### 4. Update an Item by ID

- **URL**: `/api/v1/item/<id>`
- **Method**: `PATCH`
- **Response**:
  JSON object with updated item details.

### 5. Delete an Item by ID

- **URL**: `/api/v1/item/<id>`
- **Method**: `DELETE`
- **Response**:
  `200 OK if the item is delete.`\
  `404 Not Found if the item does not exist.`\

## Setup and Usage

### Prerequisites

- Go installed

## Installation Steps

### Clone the repository:

`git clone <repository-url>`
`cd <repository-directory>`

Install required dependencies:
`go get -u github.com/gofiber/fiber/v2`
Run the application:
`go run main.go`
The API will be available at:
`http://localhost:8080`

## Thread Safety

The API uses a lock to ensure thread-safe access to the in-memory items list. This prevents race conditions when multiple threads access or modify the database concurrently.

## Notes

In-memory database: The data is not persisted; it is cleared each time the server restarts.
