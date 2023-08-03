package db

import "github.com/deepg7/go-gin-gorm/models"

func GetBookByID(id int) (models.Book, error) {
	var book models.Book
	err := models.DB.Where("id = ?", id).First(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func CreateBook(input models.CreateBookInput) models.Book {
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)
	return book
}

func GetAllBooks() []models.Book {
	var books []models.Book
	models.DB.Find(&books)
	return books
}

func UpdateBook(book models.Book, input models.UpdateBookInput) models.Book {
	models.DB.Model(&book).Updates(input)
	return book
}
