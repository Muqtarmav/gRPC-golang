package main

import (
	"coffee-shop/coffee"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("error:: %v", err)
	}
	log.Println("message:: listening on port 9000")
	s := coffee.Server{}
	grpcServer := grpc.NewServer()
	coffee.RegisterCoffeeServiceServer(grpcServer, &s)

	coffee.RegisterCoffeeUserServicesServer(grpcServer, &s)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("error:: %v", err)
	}
}
