package protos

import (
	"context"

	"orami.com/techshare/sdk"
)

type Server struct {
}

func (s *Server) GetPeople(c context.Context, pr *sdk.PeopleRequest) (*sdk.PeopleResponse, error) {

}
