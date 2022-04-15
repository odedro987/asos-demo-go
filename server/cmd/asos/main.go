package main

import (
	"asos-api/pkg/db"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	DefaultPort = "3001"
)

var dao = db.AsosDAO{}

func init() {
	dao.Server = "localhost"
	dao.Database = "asos-demo"

	dao.Connect()

	fmt.Println("connecting to DB")
}

func main() {
	r := mux.NewRouter()
	fmt.Printf("Starting server at port 3001\n")
	if err := http.ListenAndServe(":"+DefaultPort, r); err != nil {
		log.Fatal(err)
	}
}
