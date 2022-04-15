package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	DefaultPort = "3001"
)

func main() {
	r := mux.NewRouter()
	fmt.Printf("Starting server at port 3001\n")
	if err := http.ListenAndServe(":"+DefaultPort, r); err != nil {
		log.Fatal(err)
	}
}
