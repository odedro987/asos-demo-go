package handler

import (
	"asos-api/pkg/db"

	"github.com/gorilla/mux"
)

var daoInstance *db.AsosDAO

func CreateHandler(r *mux.Router, dao *db.AsosDAO) {
	mainRoutes(r)
	daoInstance = dao
}

func mainRoutes(r *mux.Router) {
	CreateItemHandler(r)
	CreateCategoriesHandler(r)
}
