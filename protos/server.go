package protos

import (
	"context"

	"orami.com/techshare/pkg"
	"orami.com/techshare/repository"
)

type Server struct {
}

func (s Server) GetPerson(c context.Context, r *pkg.PersonRequest) (*pkg.Person, error) {
	pr := repository.NewPeople()
	sp := pkg.Person{}

	person, err := pr.FindByID(r.GetId())
	if err != nil {
		return &sp, err
	}

	return &pkg.Person{
		Id:   person.ID,
		Name: person.Name,
		Address: &pkg.Address{
			City:    person.Address.City,
			Country: person.Address.Country,
		},
	}, nil
}

func (s Server) ListPeople(r *pkg.PeopleRequest, stream pkg.TechShare_ListPeopleServer) error {
	pr := repository.NewPeople()
	people, err := pr.FindAll(1000000)
	if err != nil {
		return err
	}

	for _, person := range people {
		sp := pkg.Person{
			Id:   person.ID,
			Name: person.Name,
			Address: &pkg.Address{
				City:    person.Address.City,
				Country: person.Address.Country,
			},
		}
		if err := stream.Send(&sp); err != nil {
			return err
		}
	}
	return nil
}
