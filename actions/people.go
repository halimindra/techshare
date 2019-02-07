package actions

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"orami.com/techshare/repository"
)

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var id int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("Invalid ID")
	}

	pr := repository.NewPeople()
	person := pr.FindByID(id)
	json.NewEncoder(w).Encode(person)
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	pr := repository.NewPeople()
	people := pr.FindAll(1000000)
	json.NewEncoder(w).Encode(people)
}
