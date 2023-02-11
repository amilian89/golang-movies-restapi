package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/go-pg/pg"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/movies", createMovieHandler).Methods("POST")
	r.HandleFunc("/movies", getMoviesHandler).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieHandler).Methods("GET")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies", updateMovie).Methods("PATCH")

	db := connect()
	defer db.Close()

	// Starting Server
	log.Fatal(http.ListenAndServe(":8000", r))
}

// CREATE MOVIES
func createMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get connection
	db := connect()
	defer db.Close()

	// Creating product instance
	movies := &Movie{
		ID: uuid.New().String(),
	}

	// Decoding request
	_ = json.NewDecoder(r.Body).Decode(&movies)

	// Inserting into Database
	_, err := db.Model(movies).Insert()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	// Returning movies data
	json.NewEncoder(w).Encode(movies)

}

// GET MOVIES
func getMoviesHandler(w http.ResponseWriter, r *http.Request) {
	// Load the template file
	tmpl, err := template.ParseFiles("movie.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get connection
	db := connect()
	defer db.Close()

	// Creating movies slice
	var movies []Movie
	if err := db.Model(&movies).Select(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Execute the template and write the result to the response writer
	if err := tmpl.Execute(w, movies); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GET MOVIE BY ID
func getMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get connection
	db := connect()
	defer db.Close()

	// Get ID
	params := mux.Vars(r)
	movieId := params["id"]

	movies := &Movie{ID: movieId}
	if err := db.Model(movies).WherePK().Select(); err != nil {
		log.Println(err)
		if err == pg.ErrNoRows {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No movie with that ID found"))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	// Returning movies
	json.NewEncoder(w).Encode(movies)
}

// DELETE MOVIES
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get Connect
	db := connect()
	defer db.Close()

	// Get ID
	params := mux.Vars(r)
	movieId := params["id"]

	// Creating Movie Instance
	movies := &Movie{}
	result, err := db.Model(movies).Where("id = ?", movieId).Delete()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if no rows affected
	if result.RowsAffected() == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("No movie with that ID exists")
		return
	}

	// Returning result if the delete was successful
	if result.RowsAffected() > 0 {
		json.NewEncoder(w).Encode("Movie deleted correctly")
	}

	// Returning result
	json.NewEncoder(w).Encode(result)
}

// UPDATE MOVIES
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get Connect
	db := connect()
	defer db.Close()

	// Get ID
	params := mux.Vars(r)
	movieId := params["id"]

	// Creating Actor Instance
	movies := &Movie{ID: movieId}

	_ = json.NewDecoder(r.Body).Decode(&movies)

	result, err := db.Model(movies).WherePK().Set("title = ?, year = ?, genre = ?, actor = ?, rating = ?", movies.Title, movies.Year, movies.Genre, movies.Actor, movies.Rating).Update()
	if err != nil || result.RowsAffected() == 0 {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No movie with that ID found to update"))
		return
	}

	// Returning actors
	json.NewEncoder(w).Encode(movies)
}
