package main

import (
	"github.com/zanefinner-projects/library/index"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/zanefinner-projects/library/books"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index.Handler)
	r.HandleFunc("/book", books.Create).Methods("POST")
	r.HandleFunc("/books", books.ShowAll).Methods("GET")
	r.HandleFunc("/book/", books.Delete).Methods("DELETE")
	r.HandleFunc("/book/", books.Edit).Methods("PUT")
	r.HandleFunc("/book/", books.Show).Methods("GET")
	http.ListenAndServe(":8080", r)
}