package dao

import (
	"fmt"

	"github.com/atultherajput/go_crash_course/models"
)

func (h Handler) Update(id string, updatedBook models.Book) {

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.Desc = updatedBook.Desc

	h.DB.Save(&book)

}
