package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client *mongo.Client
}

func New(ctx context.Context, conURL string) (*MongoDB, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(conURL).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.PrimaryPreferred()); err != nil {
		return nil, err
	}

	return &MongoDB{client}, nil
}

func (db *MongoDB) Close(ctx context.Context) error {
	if db.Client == nil {
		return nil
	}

	err := db.Client.Disconnect(ctx)
	if err != nil {
		return err
	}

	return nil
}

type MongoCol struct {
	Col *mongo.Collection
}

func (db *MongoDB) GetCollecion(dbName, colName string) *MongoCol {
	return &MongoCol{db.Client.Database(dbName).Collection(colName)}
}
