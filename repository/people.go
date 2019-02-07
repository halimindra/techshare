package repository

import (
	"strconv"

	"orami.com/techshare/models"
)

type People struct {
}

func NewPeople() People {
	return People{}
}

func (p *People) FindByID(id int) models.Person {
	person := models.Person{
		ID:   id,
		Name: "Person " + strconv.Itoa(id),
		Address: &models.Address{
			City:    "Jakarta",
			Country: "Indonesia",
		},
	}
	return person
}

func (p *People) FindAll(limit int) []models.Person {
	people := []models.Person{}

	for i := 1; i <= limit; i++ {
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
	return people
}
