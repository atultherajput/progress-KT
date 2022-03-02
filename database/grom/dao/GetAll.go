package dao

import (
	"fmt"

	"github.com/atultherajput/go_crash_course/models"
)

func (h Handler) GetAll() (books []models.Book) {

	if result := h.DB.Find(&books); result.Error != nil {
		fmt.Println(result.Error)
	}

	return books
}
