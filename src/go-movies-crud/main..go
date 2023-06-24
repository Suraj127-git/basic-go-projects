package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strcronv"

	"github.com/gorilla/mux"
)

type Moive struct {
	ID       string `json:"id"`
	Isbn     string `json:"isbn"`
	Title    string `json:"title"`
	Director string `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range params {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strcronv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(movies)
	params := mux.Vars(r)
	for index, item := range params {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break

		}
	}
	json.NewEncoder(w).Encode(movies)
	return
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(movies)
	params := mux.Vars(r)
	for index, item := range params {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strcronv.Itoa(rand.Intn(100000000))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
	json.NewEncoder(w).Encode(movies)
	return
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "213123", Title: "Movie one", Director: &Director{FirstName: "Suraj", LastName: "shetty"}})
	movies = append(movies, Movie{ID: "2", Isbn: "3323233", Title: "Movie second", Director: &Director{FirstName: "sandeep", LastName: "suthar"}})

	r.handleFunc("/movies", getMovies).Methods("GET")
	r.handleFunc("/movies/{id}", getMovie).Methods("GET")
	r.handleFunc("/movies", createMovie).Methods("POST")
	r.handleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.handleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
