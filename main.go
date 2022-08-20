package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// define movie struct
type Movie struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Genre    string    `json:"genre"`
	Isbn     string    `json:"isbn"`
	Director *Director `json:"director"`
}

// define director struct
type Director struct {
	Name        string `json:"director name"`
	Nationality string `json:"nationality"`
}

// create slice of movies
var Movies []Movie

// declare handler funcs
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(Movies[index])
			return
		}
	}
	fmt.Fprintf(w, "Movie with that id not found")
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		fmt.Fprintf(w, "wrong method associated with the request")
	}
	var newMovie Movie
	json.NewDecoder(r.Body).Decode(&newMovie)
	newMovie.Id = strconv.Itoa(rand.Intn(10000))
	Movies = append(Movies, newMovie)
	json.NewEncoder(w).Encode(Movies)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Movies {
		if item.Id == params["id"] {
			Movies = append(Movies[:index], Movies[(index+1):]...)
		}
	}
	var updatedMovie Movie
	json.NewDecoder(r.Body).Decode(&updatedMovie)
	updatedMovie.Id = params["id"]
	Movies = append(Movies, updatedMovie)
	json.NewEncoder(w).Encode(Movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Movies {
		if item.Id == params["id"] {
			Movies = append(Movies[:index], Movies[(index+1):]...)
		}
	}
}

func main() {

	//add movies to the slice of movies(necessary because we are not using a database)
	Movies = append(Movies, Movie{Id: "1", Name: "first movie", Genre: "Action", Isbn: "101010", Director: &Director{Name: "Tolu Madamori", Nationality: "Nigerian"}})

	//create a new router
	r := mux.NewRouter()

	//create handlers
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	//start server
	fmt.Printf("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Panic("failed to start server on that port")
	}
}
