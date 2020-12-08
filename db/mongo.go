package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	//Client ...
	clientInstance *mongo.Client
	err            error
	mongoOnce      sync.Once
)

const (
	uri = "mongodb://root:password@localhost:27017/"
)

//GetConnect ...
func GetClient() *mongo.Client {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(uri)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			panic(err)
		}
		err = client.Ping(context.TODO(), nil)

		if err != nil {
			panic("Ping error")
		}
		clientInstance = client
	})
	return clientInstance
}
