package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	ItemPrefix = "/items"
)

func CreateItemHandler(r *mux.Router) {
	mainItemRoutes(r)
}

func mainItemRoutes(r *mux.Router) {
	s := r.PathPrefix(ItemPrefix).Subrouter()

	s.HandleFunc("/", findAllItems).Methods("GET")
	s.HandleFunc("/{id}", findItemById).Methods("GET")
	s.HandleFunc("/{brand}", findItemsByBrand).Methods("GET")
	s.HandleFunc("/{category}", findItemsByCategory).Methods("GET")
	s.HandleFunc("/{gender}", findItemsByGender).Methods("GET")
	s.HandleFunc("/{gender}/{category}", findItemsByGenderAndCategory).Methods("GET")
}

func findAllItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func findItemById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func findItemsByBrand(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func findItemsByCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func findItemsByGender(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func findItemsByGenderAndCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}
