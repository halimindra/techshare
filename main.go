package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"orami.com/techshare/actions"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Invalid argument")
	}

	switch os.Args[1] {
	case "rest":
		router := mux.NewRouter()
		router.HandleFunc("/people", actions.GetPeople).Methods("GET")
		router.HandleFunc("/people/{id}", actions.GetPerson).Methods("GET")
		log.Fatal(http.ListenAndServe(":8000", router))
	case "grpc":
		log.Println("Starting gRPC Server")
	default:
		log.Fatal("Invalid argument")
	}
}
