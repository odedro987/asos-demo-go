package handler

import (
	utils "asos-api/pkg/json"
	"asos-api/pkg/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

const (
	CategoriesPrefix = "/categories"
)

func CreateCategoriesHandler(r *mux.Router) {
	mainCategoryRoutes(r)
}

func mainCategoryRoutes(r *mux.Router) {
	s := r.PathPrefix(CategoriesPrefix).Subrouter()

	s.HandleFunc("", findAll).Methods("GET")
	s.HandleFunc("", addCategory).Methods("POST")
}

func addCategory(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var category model.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	category.ID = bson.NewObjectId()
	if err := daoInstance.InsertCategory(category); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusCreated, category)
}

func findAll(w http.ResponseWriter, r *http.Request) {
	categories, err := daoInstance.FindAllCategories()

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJson(w, http.StatusOK, categories)
}
