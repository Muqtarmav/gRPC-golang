package coffee

import (
	"coffee-shop/utils"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type User struct {
	Id              primitive.ObjectID
	Name            string
	EmailAddress    string
	DeliveryAddress string
	PhoneNumber     string
	Password        string
}

var usersCollection *mongo.Collection
var client = utils.Connect()

func (s *Server) Register(ctx context.Context, request *RegisterRequest) (*Response, error) {
	//TODO implement me
	log.Println(client.Ping(context.TODO(), nil))
	usersCollection = client.Database("coffee_user").Collection("user")
	if _, err := usersCollection.InsertOne(ctx, bson.M{
		"Name":            request.GetName(),
		"EmailAddress":    request.GetEmailAddress(),
		"DeliveryAddress": request.GetDeliveryAddress(),
		"PhoneNumber":     request.GetPhoneNumber(),
		"Password":        request.GetPassword(),
	}); err != nil {
		log.Fatal(errors.New("Failed to register user"))
	}
	response := &Response{
		Message: "USer successfully registered",
	}
	return response, nil

}

func (s *Server) Login(ctx context.Context, request *LoginRequest) (*Response, error) {
	//TODO implement me
	var user = &User{}
	ctxs, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	err := usersCollection.FindOne(ctxs, bson.M{"EmailAddress": request.Email}).Decode(user)
	if err != nil {
		log.Fatalln(err)
	}
	err = usersCollection.FindOne(ctxs, bson.M{"Password": request.Password}).Decode(user)
	if err != nil {
		log.Fatalln(err)
	}
	response := &Response{
		Message: "Successfully logged in",
	}
	return response, nil

}

func (s *Server) AddCo(ctx context.Context, cart *RegisterRequest) (*Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) mustEmbedUnimplementedCoffeeUserServicesServer() {
	//TODO implement me
	panic("implement me")
}
