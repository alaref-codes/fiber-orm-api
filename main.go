package main

import (
	"fmt"
	"log"

	"github.com/alaref-codes/fiber-app/book"
	"github.com/alaref-codes/fiber-app/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/book", book.GetBooks)
	app.Get("/api/book/:id", book.GetBook)
	app.Post("/api/book", book.PostBooks)
	app.Delete("/api/book/:id", book.DeleteBooks)

}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection open ed")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Print("Database migrated ")
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
