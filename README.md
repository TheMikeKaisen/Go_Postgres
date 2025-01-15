# Go Fiber + PostgreSQL Project

This project is a simple CRUD (Create, Read, Update, Delete) application built using the Go programming language, the Fiber web framework, and PostgreSQL as the database. The project demonstrates fundamental concepts for building a backend with Go, including:

- Handling HTTP requests using Fiber.
- Integrating PostgreSQL with GORM.
- Structuring a Go project with models, routes, and database configurations.

---

## Features

1. **Create a Book**
   - Endpoint: `POST /api/createBook`
   - Creates a new book entry in the database.

2. **Get All Books**
   - Endpoint: `GET /api/books`
   - Retrieves all book entries from the database.

3. **Get Book by ID**
   - Endpoint: `GET /api/getBook/:id`
   - Retrieves a single book by its ID.

4. **Delete a Book by ID**
   - Endpoint: `DELETE /api/deleteBook/:id`
   - Deletes a book entry from the database based on its ID.

---

## Technologies Used

### 1. **Go**
   - Go is used as the primary programming language for this project.

### 2. **Fiber**
   - Fiber is a lightweight and fast web framework for Go.
   - Official documentation: [Fiber Docs](https://gofiber.io/)

### 3. **PostgreSQL**
   - PostgreSQL is the database used to store book records.

### 4. **GORM**
   - GORM is used as the ORM (Object Relational Mapper) to interact with the PostgreSQL database.
   - Official documentation: [GORM Docs](https://gorm.io/)

---

## Prerequisites

Before running this project, ensure you have the following installed:

1. [Go](https://golang.org/) (version 1.19 or higher)
2. [PostgreSQL](https://www.postgresql.org/)
3. [Git](https://git-scm.com/)

---

## Installation and Setup

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/TheMikeKaisen/Go_Postgres.git
   cd your-repo-name
   ```

2. **Install Dependencies:**
   ```bash
   go mod tidy
   ```

3. **Set Up the Environment File:**
   Create a `.env` file in the root directory and add your database configuration:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=your_database_name
   DB_SSLMODE=disable
   ```

4. **Run the PostgreSQL Server:**
   Ensure your PostgreSQL server is running.

5. **Run the Application:**
   ```bash
   go run main.go
   ```

6. **Access the Application:**
   The application runs on `http://localhost:8080` by default.

---

## Project Structure

```plaintext
.
├── main.go             # Application entry point
├── models.go           # GORM model definitions
├── postgres.go         # PostgreSQL database connection setup
├── go.mod              # Go module dependencies
├── go.sum              # Go module checksum
├── .env                # Environment variables for database configuration
└── README.md           # Project documentation
```

---

## API Endpoints

### 1. **Create a Book**
   - **Endpoint:** `POST /api/createBook`
   - **Request Body:**
     ```json
     {
       "author": "Author Name",
       "name": "Book Title",
       "publisher": "Publisher Name"
     }
     ```

### 2. **Get All Books**
   - **Endpoint:** `GET /api/books`
   - **Response:**
     ```json
     {
       "message": "Books retrieved successfully",
       "data": [
         {
           "id": 1,
           "author": "Author Name",
           "name": "Book Title",
           "publisher": "Publisher Name"
         }
       ]
     }
     ```

### 3. **Get Book by ID**
   - **Endpoint:** `GET /api/getBook/:id`
   - **Response:**
     ```json
     {
       "message": "Found your book!",
       "data": {
         "id": 1,
         "author": "Author Name",
         "name": "Book Title",
         "publisher": "Publisher Name"
       }
     }
     ```

### 4. **Delete a Book by ID**
   - **Endpoint:** `DELETE /api/deleteBook/:id`
   - **Response:**
     ```json
     {
       "message": "Book deleted successfully!"
     }
     ```

---

## Key Concepts Learned

### 1. **Go (Golang)**
   - Structs for modeling data.
   - Fiber for handling HTTP routes and responses.
   - Error handling in Go.

### 2. **Fiber**
   - Routing: `GET`, `POST`, and `DELETE` routes.
   - Parsing JSON requests and sending JSON responses.

### 3. **PostgreSQL**
   - Basic configuration and connection setup.
   - Auto-migration of models using GORM.
   - Performing CRUD operations with GORM.

---

## Improvements to Consider

1. Add middleware for authentication and validation.
2. Implement proper error logging and monitoring.
3. Extend the project to include updates (PUT endpoint).
4. Write unit tests for each route and database interaction.

---

