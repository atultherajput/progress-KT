package gin_controller

import (
	"net/http"

	"github.com/atultherajput/go_crash_course/database/grom/dao"
	"github.com/atultherajput/go_crash_course/models"
	"github.com/gin-gonic/gin"
)

var Handler dao.Handler

func AddBook(c *gin.Context) {
	// Read to request body
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Append to the Books table
	Handler.Add(book)

	// Send a 201 created response
	c.JSON(http.StatusCreated, gin.H{"data": "The book has been inserted successfully!"})
}

func GetAllBooks(c *gin.Context) {

	//Find all books
	books := Handler.GetAll()

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func GetBook(c *gin.Context) {
	// Read dynamic id parameter
	id := c.Param("id")

	// Find book by Id
	book := Handler.Get(id)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	// Read dynamic id parameter
	id := c.Param("id")

	// Read request body
	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Update DB
	Handler.Update(id, updatedBook)

	c.JSON(http.StatusOK, gin.H{"data": "The book has been upadted successfully!"})
}

func DeleteBook(c *gin.Context) {
	// Read the dynamic id parameter
	id := c.Param("id")

	// Delete that book
	Handler.Delete(id)

	c.JSON(http.StatusOK, gin.H{"data": "The book has been deleted successfully!"})
}
