package database

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"yas/cfg"
)

var clientInstance *mongo.Client
var clientError error
var once sync.Once

func GetMongoClient(ctx context.Context) (*mongo.Client, error) {
	once.Do(func() {
		clientOptions := options.Client().ApplyURI(cfg.MongoDbURI())
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			clientError = err
		}

		if err = client.Ping(ctx, nil); err != nil {
			clientError = err
		}
		clientInstance = client
	})
	return clientInstance, clientError
}
