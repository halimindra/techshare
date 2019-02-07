package actions

import (
	"encoding/json"
	"net/http"
	"strconv"

	"orami.com/techshare/models"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	people := []models.Person{}

	for i := 1; i <= 10000; i++ {
		people = append(
			people,
			models.Person{
				ID:   i,
				Name: "Person " + strconv.Itoa(i),
				Address: &models.Address{
					City:    "Jakarta",
					Country: "Indonesia",
				},
			},
		)
	}

	json.NewEncoder(w).Encode(people)
}
