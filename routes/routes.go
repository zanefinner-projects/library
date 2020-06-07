package routes

import "github.com/gorilla/mux"

//CreateRoutes generates routes for gorilla
func CreateRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", index.Handler)
}
