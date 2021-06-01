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

//Books model
type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:author`
}

type Author struct {
	Fistname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var books []Book

//Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
	fmt.Println("Books sent")
}

//Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var vars = mux.Vars(r)
	for _, bok := range books {
		if bok.ID == vars["id"] {
			json.NewEncoder(w).Encode(bok)
			break
		}
	}
}

//Create a  books
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//update a  books
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var vars = mux.Vars(r)
	var book Book
	for index, item := range books {
		if item.ID == vars["id"] {
			_ = json.NewDecoder(r.Body).Decode(&book)
			item.Title = book.Title
			item.Author = book.Author
			books[index] = item
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

//Delete a  books
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for ind, item := range books {
		if item.ID == params["id"] {
			books = append(books[:ind], books[ind+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}

func main() {
	fmt.Println("Starting restful api...")
	router := mux.NewRouter()

	// Mock books for test
	books = append(books, Book{ID: "1", Title: "Au pays", Author: &Author{Fistname: "John", Lastname: "sidik"}})
	books = append(books, Book{ID: "2", Title: "Au pays des arbres", Author: &Author{Fistname: "Jonas", Lastname: "sidi"}})
	books = append(books, Book{ID: "5", Title: "Au pays des eaux", Author: &Author{Fistname: "Joe", Lastname: "sid"}})

	//Router
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/book", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
