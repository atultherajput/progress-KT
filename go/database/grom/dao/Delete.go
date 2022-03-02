package dao

import (
	"fmt"

	"github.com/atultherajput/go_crash_course/models"
)

func (h Handler) Delete(id string) {

	// Find the book by Id
	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Delete that book
	h.DB.Delete(&book)

}
