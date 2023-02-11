package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

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

	r.HandleFunc("/actors", createActorHandler).Methods("POST")
	r.HandleFunc("/actors", getActorsHandler).Methods("GET")
	r.HandleFunc("/actors/{id}", getActorHandler).Methods("GET")
	r.HandleFunc("/actors/{id}", deleteActor).Methods("DELETE")
	r.HandleFunc("/actors", updateActor).Methods("PATCH")

	db := connect()
	defer db.Close()

	// Starting Server
	log.Fatal(http.ListenAndServe(":8000", r))
}

// CREATE ACTOR
func createActorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get connection
	db := connect()
	defer db.Close()

	// Creating actor instance
	actors := &Actor{
		ID: uuid.New().String(),
	}

	// Decoding request
	_ = json.NewDecoder(r.Body).Decode(&actors)

	// Inserting into Database
	_, err := db.Model(actors).Insert()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	// Returning actors data
	json.NewEncoder(w).Encode(actors)

}

// GET ACTORS
func getActorsHandler(w http.ResponseWriter, r *http.Request) {
	// Load the template file
	tmpl, err := template.ParseFiles("actor.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get connection
	db := connect()
	defer db.Close()

	// Creating movies slice
	var actor []Actor
	if err := db.Model(&actor).Select(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Execute the template and write the result
	if err := tmpl.Execute(w, actor); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GET ACTOR BY ID
func getActorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get connection
	db := connect()
	defer db.Close()

	// Get ID
	params := mux.Vars(r)
	actorId := params["id"]

	actors := &Actor{ID: actorId}
	if err := db.Model(actors).WherePK().Select(); err != nil {
		log.Println(err)
		if err == pg.ErrNoRows {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No actor with that ID found"))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	// Returning movies
	json.NewEncoder(w).Encode(actors)
}

// DELETE ACTOR BY ID
func deleteActor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get Connect
	db := connect()
	defer db.Close()

	// Get ID
	params := mux.Vars(r)
	actorId := params["id"]

	// Creating Actor Instance
	actors := &Actor{}
	result, err := db.Model(actors).Where("id = ?", actorId).Delete()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if no rows affected
	if result.RowsAffected() == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("No actor with that ID exists")
		return
	}

	// Returning result if the delete was successful
	if result.RowsAffected() > 0 {
		json.NewEncoder(w).Encode("Actor deleted correctly")
	}

	// Returning result
	json.NewEncoder(w).Encode(result)
}

// UPDATE ACTOR
func updateActor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get Connect
	db := connect()
	defer db.Close()

	// Get ID
	params := mux.Vars(r)
	actorId := params["id"]

	// Creating Actor Instance
	actors := &Actor{ID: actorId}

	_ = json.NewDecoder(r.Body).Decode(&actors)

	result, err := db.Model(actors).WherePK().Set("first_name = ?, last_name = ?, gender = ?, age = ?, movies = ?", actors.FirstName, actors.LastName, actors.Gender, actors.Age, actors.Movies).Update()
	if err != nil || result.RowsAffected() == 0 {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No actor with that ID found to update"))
		return
	}

	// Returning actors
	json.NewEncoder(w).Encode(actors)
}
