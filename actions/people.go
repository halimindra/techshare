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
	person, err := pr.FindByID(int64(id))
	if err != nil {
		log.Print(err)
	}

	json.NewEncoder(w).Encode(person)
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	pr := repository.NewPeople()
	people, err := pr.FindAll(1000000)
	if err != nil {
		log.Print(err)
	}

	json.NewEncoder(w).Encode(people)
}
