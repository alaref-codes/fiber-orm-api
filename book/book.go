package book

import (
	"github.com/alaref-codes/fiber-app/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

/*
 The Ctx contains is short for context it contains all the attribute
 That come with a http request
 and contains the method we need to complete a specific response
*/

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)

}

func GetBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		// fiber.NewError(fiber. , err)
		return err
		// return fiber.ErrNotFound
	}
	db := database.DBConn
	var book Book

	if err := db.Find(&book, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return fiber.ErrNotFound
		}
		return err
	}

	// db.Find(&book, id)
	// if book.ID == 0 {
	// 	return fiber.ErrNotFound
	// }
	return c.JSON(book)
}

func PostBooks(c *fiber.Ctx) error {
	db := database.DBConn
	// book := &Book{}  // Both are the same
	book := new(Book)
	// var book Book
	err := c.BodyParser(book)
	if err != nil {
		return err
	}

	db.Create(&book)
	return c.JSON(book)
}

func DeleteBooks(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		// fiber.NewError(fiber. , err)
		return err
		// return fiber.ErrNotFound
	}
	db := database.DBConn
	var book Book

	db.First(&book, id)
	if book.Title == "" {
		return c.Status(500).SendString("No book was found brother !")
	}

	db.Delete(&book)
	return c.SendString("Deleted successfully")
}
