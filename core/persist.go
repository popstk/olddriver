package core

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mgoDB  = "spider"
	mgoURL = "mongodb://127.0.0.1:27017/"
)

// Collection -
func Collection(key string) (*mongo.Collection, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mgoURL))
	Must(err)

	c := client.Database(mgoDB).Collection(key)
	return c, nil
}
