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

func (p *People) FindByID(id int64) (models.Person, error) {
	person := models.Person{
		ID:   id,
		Name: "Person " + strconv.Itoa(int(id)),
		Address: &models.Address{
			City:    "Jakarta",
			Country: "Indonesia",
		},
	}
	return person, nil
}

func (p *People) FindAll(limit int64) ([]models.Person, error) {
	people := []models.Person{}

	for i := 1; i <= int(limit); i++ {
		people = append(
			people,
			models.Person{
				ID:   int64(i),
				Name: "Person " + strconv.Itoa(i),
				Address: &models.Address{
					City:    "Jakarta",
					Country: "Indonesia",
				},
			},
		)
	}
	return people, nil
}
