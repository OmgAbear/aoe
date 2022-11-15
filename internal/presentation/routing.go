package presentation

import (
	"github.com/gorilla/mux"
	"net/http"
)

//RegisterRoutes sets the handlers for a specific route
func RegisterRoutes(router *mux.Router) {
	superheroHandler := NewSuperheroHandler()
	router.HandleFunc("/api/v1/superhero", superheroHandler.List).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/superhero", superheroHandler.Create).Methods(http.MethodPost)
}
