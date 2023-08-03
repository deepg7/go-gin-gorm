package controllers

import (
	"net/http"
	"strconv"

	"github.com/deepg7/go-gin-gorm/db"
	"github.com/deepg7/go-gin-gorm/models"

	"github.com/gin-gonic/gin"
)

// GetBooks         godoc
// @Summary      Get all book
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         books
// @Produce      json
// @Success      200   {object}  models.Book
// @Router       /book [get]
func FindBooks(c *gin.Context) {
	books := db.GetAllBooks()
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// CreateBook             godoc
// @Summary      Store a new book
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         books
// @Produce      json
// @Param        book  body      models.Book  true  "Book JSON"
// @Success      200   {object}  models.Book
// @Router       /book [post]
func CreateBook(c *gin.Context) {
	// Validate input
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := db.CreateBook(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBook(c *gin.Context) {
	var book models.Book
	id, er := strconv.Atoi(c.Param("id"))
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}
	book, err := db.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {

	id, er := strconv.Atoi(c.Param("id"))
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}
	book, err := db.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBook := db.UpdateBook(book, input)
	c.JSON(http.StatusOK, gin.H{"data": updatedBook})
}
