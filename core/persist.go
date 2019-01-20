package core

import (
	"github.com/mongodb/mongo-go-driver/mongo"
)

const (
	mgoDB  = "spider"
	mgoURL = "mongodb://127.0.0.1:27017/"
)

// Collection -
func Collection(key string) (*mongo.Collection, error) {
	client, err := mongo.NewClient(mgoURL)
	if err != nil {
		return nil, err
	}

	err = client.Connect(nil)
	if err != nil {
		return nil, err
	}

	c := client.Database(mgoDB).Collection(key)
	return c, nil
}
