package net

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atultherajput/go_crash_course/database/grom/dao"
	"github.com/atultherajput/go_crash_course/database/grom/dbinit"
	mux_controller "github.com/atultherajput/go_crash_course/net/controllers/mux"
	"github.com/gorilla/mux"
)

func RunMux(port *int) {
	DB := dbinit.Init()
	mux_controller.Handler = dao.New(DB)

	router := mux.NewRouter()

	//Middleware Intercepter
	router.Use(muxCustomMiddleware)

	router.HandleFunc("/books", mux_controller.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", mux_controller.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", mux_controller.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", mux_controller.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", mux_controller.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(fmt.Sprintf(":%d", *port), router)
}
