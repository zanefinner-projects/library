package main

import (
	"github.com/zanefinner-projects/library/index"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index.Handler)
	http.ListenAndServe(":8080", r)
}