package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"orami.com/techshare/actions"
	"orami.com/techshare/protos"
	pb "orami.com/techshare/pkg"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Invalid argument")
	}

	switch os.Args[1] {
	case "rest":
		log.Println("Starting Rest API Server")

		router := mux.NewRouter()
		router.HandleFunc("/people", actions.GetPeople).Methods("GET")
		router.HandleFunc("/people/{id}", actions.GetPerson).Methods("GET")

		log.Fatal(http.ListenAndServe(":8000", router))
	case "grpc":
		log.Println("Starting gRPC Server")
		flag.Parse()

		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		techShareServer := protos.Server{}
		pb.RegisterTechShareServer(grpcServer, techShareServer)
		grpcServer.Serve(lis)
	default:
		log.Fatal("Invalid argument")
	}
}
