package main

import (
	"log"
	"net/http"
	"os"

	"Github.com/TheMikeKasien/Go_PostGres/models"
	"Github.com/TheMikeKasien/Go_PostGres/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type Book struct {
	Author    string `json:"author"`
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	var book Book

	err := context.BodyParser(&book)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": "could not create book"})
		return err
	}

	err = r.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "error while creating a book!"})
		return err
	}

	context.Status(http.StatusOK).JSON(fiber.Map{"message": "book created successfully", "data": book})
	return nil

}

func (r *Repository) GetAllBooks(context *fiber.Ctx) error {
	books := []models.Books{}

	err := r.DB.Find(&books).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Could not find books!"})
		return err
	}

	context.Status(http.StatusOK).JSON(fiber.Map{"message": "Books got successfully", "data": books})
	return nil
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	// extract the id
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "id field is empty"})
		return nil
	}

	var bookModel models.Books
	err := r.DB.Delete(bookModel, id).Error // it takes reference from the model's structure and looks inside table where primary key == id and then delete it.

	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "error while deleting the book"})
		return err
	}

	context.Status(http.StatusOK).JSON(fiber.Map{"message": "Book deleted Successfully!"})
	return nil

}

func (r *Repository) GetBookById(context *fiber.Ctx) error {
	// extract the id from url
	id := context.Params("id")
	if id == "" {
		context.Status(400).JSON(fiber.Map{"message": "id field is empty"})
		return nil
	}

	book := &models.Books{}
	err := r.DB.Find(book, id).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "error while finding book."})
		return err
	}

	context.Status(200).JSON(fiber.Map{"message": "Found your book!", "data": book})
	return nil

}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/getBook/:id", r.GetBookById)
	api.Post("/createBook", r.CreateBook)
	api.Delete("/deleteBook/:id", r.DeleteBook)
	api.Get("/books", r.GetAllBooks)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// created a server
	app := fiber.New()

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatalf("could not migrate db")
	}

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatalf("could not migrate db")
	}

	r := Repository{
		DB: db,
	}

	r.SetupRoutes(app)

	app.Listen(":8080")

}
