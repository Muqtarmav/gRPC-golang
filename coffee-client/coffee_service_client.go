package main

import (
	"coffee-shop/coffee"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial("34.170.197.145:6000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("error dialing server::", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("error closing client conn::", err)

		}
	}(conn)
	coffeeService := coffee.NewCoffeeServiceClient(conn)
	response, err := coffeeService.AddCoffee(context.TODO(), &coffee.AddCoffeeRequest{
		Name:  "coffee black",
		Price: 30.00,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)

	userService := coffee.NewCoffeeUserServicesClient(conn)
	addUseresponse, err := userService.Register(context.TODO(), &coffee.RegisterRequest{
		Name:         "John G",
		PhoneNumber:  "0099990",
		EmailAddress: "Johndoe@gmail.com",
		Password:     "DoeJohn",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(addUseresponse)

	loginUser, err := userService.Login(context.TODO(), &coffee.LoginRequest{
		Email:    "Johndoe@gmail.com",
		Password: "DoeJohn",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(loginUser)
}
