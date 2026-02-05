package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Director *Director `json:"director"`
	Year uint `json:"year"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

// GET: get list of all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// PATCH: create a new movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie

	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.IntN(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}


func main() {
	r := mux.NewRouter()
	// TO DO - homepage 
	r.HandleFunc("movies", getMovies).Methods("GET")
	// r.HandleFunc("movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("movies", createMovie).Methods("POST")
	// r.HandleFunc("movies/{id}", updateMovie).Methods("PUT")
	// r.HandleFunc("movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}