package net

import (
	"fmt"
	"net/http"

	"github.com/atultherajput/go_crash_course/database/grom/dao"
	"github.com/atultherajput/go_crash_course/database/grom/dbinit"
	gin_controller "github.com/atultherajput/go_crash_course/net/controllers/gin"
	"github.com/gin-gonic/gin"
)

func RunGin(port *int) {
	DB := dbinit.Init()
	gin_controller.Handler = dao.New(DB)

	router := gin.Default()
	router.Use(ginCustomMiddleware)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	router.GET("/books", gin_controller.GetAllBooks)
	router.GET("/books/:id", gin_controller.GetBook)
	router.POST("/books", gin_controller.AddBook)
	router.PUT("/books/:id", gin_controller.UpdateBook)
	router.DELETE("/books/:id", gin_controller.DeleteBook)

	router.Run(fmt.Sprintf(":%d", *port))
}
