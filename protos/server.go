package protos

import (
	"context"

	"orami.com/techshare/repository"
	"orami.com/techshare/sdk"
)

type Server struct {
}

func (s *Server) GetPerson(c context.Context, r *sdk.PersonRequest) (*sdk.Person, error) {
	pr := repository.NewPeople()
	sp := sdk.Person{}

	person, err := pr.FindByID(r.GetId())
	if err != nil {
		return &sp, err
	}

	return &sdk.Person{
		Id:   person.ID,
		Name: person.Name,
		Address: &sdk.Address{
			City:    person.Address.City,
			Country: person.Address.Country,
		},
	}, nil
}

func (s *Server) ListPeople(r *sdk.PeopleRequest, stream sdk.TechShare_ListPeopleServer) error {
	pr := repository.NewPeople()
	people, err := pr.FindAll(1000000)
	if err != nil {
		return err
	}

	for _, person := range people {
		sp := sdk.Person{
			Id:   person.ID,
			Name: person.Name,
			Address: sdk.Address{
				City:    person.Address.City,
				Country: person.Address.Country,
			},
		}
		if err := stream.Send(sp); err != nil {
			return err
		}
	}
	return nil
}
