package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	//Client ...
	Client *mongo.Client

	//Ctx ...
	Ctx context.Context
	err error
)

//Connect ...
func Connect(URI string) {
	Client, err = mongo.NewClient(options.Client().ApplyURI(URI))

	if err != nil {
		panic("Failed to make new client!")
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = Client.Connect(ctx)

	if err != nil {
		panic("Cannot connect to db!")
	}
}
