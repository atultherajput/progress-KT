package net

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/atultherajput/go_crash_course/database"
	"github.com/atultherajput/go_crash_course/models"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Movie struct {
	MovieID   string `json:"movieid"`
	MovieName string `json:"moviename"`
}

// DB set up
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", database.DB_HOST, database.DB_PORT, database.DB_USER, database.DB_PASSWORD, database.DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

// Function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Get all movies
// response and request handlers
func GetMovies(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Getting movies...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM movies")

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var movies []Movie

	// Foreach movie
	for rows.Next() {
		var id int
		var movieID string
		var movieName string

		err = rows.Scan(&id, &movieID, &movieName)

		// check errors
		checkErr(err)

		movies = append(movies, Movie{MovieID: movieID, MovieName: movieName})
	}

	var response = models.JsonResponse{Type: "success", Data: movies}

	json.NewEncoder(w).Encode(response)
}

// Get movie by id
// response and request handlers
func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	movieID := params["movieid"]

	db := setupDB()

	printMessage("Getting movie: " + movieID)

	// Get movie from movies table that have movieID
	rows, err := db.Query("SELECT * FROM movies where movieID = $1;", movieID)

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var movies []Movie

	// Foreach movie
	for rows.Next() {
		var id int
		var movieID string
		var movieName string

		err = rows.Scan(&id, &movieID, &movieName)

		// check errors
		checkErr(err)

		movies = append(movies, Movie{MovieID: movieID, MovieName: movieName})
	}

	var response = models.JsonResponse{Type: "success", Data: movies}

	json.NewEncoder(w).Encode(response)
}

// Create a movie
// response and request handlers
func CreateMovie(w http.ResponseWriter, r *http.Request) {

	var movie Movie

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	movieID := movie.MovieID
	movieName := movie.MovieName

	var response = models.JsonResponse{}

	if movieID == "" || movieName == "" {
		response = models.JsonResponse{Type: "error", Message: "You are missing movieID or movieName parameter."}
	} else {
		db := setupDB()

		printMessage("Inserting movie into DB")

		fmt.Println("Inserting new movie with ID: " + movieID + " and name: " + movieName)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO movies(movieID, movieName) VALUES($1, $2) returning id;", movieID, movieName).Scan(&lastInsertID)

		// check errors
		checkErr(err)

		response = models.JsonResponse{Type: "success", Message: "The movie has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// Update a movie name
// response and request handlers
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	movieID := params["movieid"]
	movieName := r.FormValue("moviename")

	var response = models.JsonResponse{}

	if movieID == "" {
		response = models.JsonResponse{Type: "error", Message: "You are missing movieID parameter."}
	} else {
		db := setupDB()

		printMessage(fmt.Sprintf("updating movie name %s  with movieID %s from DB: ", movieName, movieID))

		sqlStatement := `UPDATE movies SET movieName = $1 WHERE movieID = $2;`
		_, err := db.Exec(sqlStatement, movieName, movieID)

		// check errors
		checkErr(err)

		response = models.JsonResponse{Type: "success", Message: "The movie has been upadted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// Delete a movie
// response and request handlers
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	movieID := params["movieid"]

	var response = models.JsonResponse{}

	if movieID == "" {
		response = models.JsonResponse{Type: "error", Message: "You are missing movieID parameter."}
	} else {
		db := setupDB()

		printMessage("Deleting movie from DB: " + movieID)

		_, err := db.Exec("DELETE FROM movies where movieID = $1", movieID)

		// check errors
		checkErr(err)

		response = models.JsonResponse{Type: "success", Message: "The movie has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// Delete all movies
// response and request handlers
func DeleteMovies(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Deleting all movies...")

	_, err := db.Exec("DELETE FROM movies")

	// check errors
	checkErr(err)

	printMessage("All movies have been deleted successfully!")

	var response = models.JsonResponse{Type: "success", Message: "All movies have been deleted successfully!"}

	json.NewEncoder(w).Encode(response)
}

// Run function
func RunMuxPq(port *int) {
	fmt.Println("PG DB Connection")
	// Init the mux router
	router := mux.NewRouter()

	//Middleware Intercepter
	router.Use(muxCustomMiddleware)

	// Route handles & endpoints

	// Get all movies
	router.HandleFunc("/movies/", GetMovies).Methods("GET")

	// Get a specific movie by the movieID
	router.HandleFunc("/movies/{movieid}", GetMovie).Methods("GET")

	// Create a movie
	router.HandleFunc("/movies/", CreateMovie).Methods("POST")

	// Update a specific movie by the movieID
	router.HandleFunc("/movies/{movieid}", UpdateMovie).Methods("PUT")

	// Delete a specific movie by the movieID
	router.HandleFunc("/movies/{movieid}", DeleteMovie).Methods("DELETE")

	// Delete all movies
	router.HandleFunc("/movies/", DeleteMovies).Methods("DELETE")

	// serve the app
	fmt.Printf("Server at %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}
