package dao

import (
	"fmt"

	"github.com/atultherajput/go_crash_course/models"
)

func (h Handler) Get(id string) (book models.Book) {

	// Find book by Id
	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	return book

}
