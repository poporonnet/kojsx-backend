package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	Cli *mongo.Client
}

func NewMongoDBClient(dsn string) *Client {
	uri := options.Client().ApplyURI(dsn)
	if err := uri.Validate(); err != nil {
		panic(err)
	}

	cli, err := mongo.Connect(context.Background(), uri)
	if err != nil {
		panic(err)
	}

	return &Client{Cli: cli}
}
