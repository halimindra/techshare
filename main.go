package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"orami.com/techshare/actions"
	pb "orami.com/techshare/pkg"
	"orami.com/techshare/proto"
)

var (
	mode = flag.String("mode", "rest", "Server mode, options: rest, grpc")
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	flag.Parse()

	switch *mode {
	case "rest":
		log.Println("Starting Rest API Server")

		router := mux.NewRouter()
		router.HandleFunc("/people", actions.GetPeople).Methods("GET")
		router.HandleFunc("/people/{id}", actions.GetPerson).Methods("GET")

		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
	case "grpc":
		log.Println("Starting gRPC Server")

		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		techShareServer := proto.Server{}
		pb.RegisterTechShareServer(grpcServer, techShareServer)
		grpcServer.Serve(lis)
	default:
		log.Fatal("Invalid argument")
	}
}
