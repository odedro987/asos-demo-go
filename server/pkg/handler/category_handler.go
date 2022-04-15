package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	CategoriesPrefix = "/categories"
)

func CreateCategoriesHandler(r *mux.Router) {
	mainCategoryRoutes(r)
}

func mainCategoryRoutes(r *mux.Router) {
	s := r.PathPrefix(CategoriesPrefix).Subrouter()

	s.HandleFunc("/", findAll).Methods("GET")
}

func findAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}
