package dao

import (
	"fmt"

	"github.com/atultherajput/go_crash_course/models"
)

func (h Handler) Add(book models.Book) {

	// Append to the Books table
	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

}
