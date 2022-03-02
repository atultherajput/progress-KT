package dbinit

import (
	"fmt"
	"log"

	"github.com/atultherajput/go_crash_course/database"
	"github.com/atultherajput/go_crash_course/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", database.DB_HOST, database.DB_PORT, database.DB_USER, database.DB_PASSWORD, database.DB_NAME, database.TIME_ZONE)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}
