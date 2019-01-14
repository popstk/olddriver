package core

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

const (
	mgoDB  = "spider"
	mgoURL = "mongodb://127.0.0.1:27017/"
)

var client *mongo.Client

// Item -
type Item struct {
	Title  string    `json:"title"`
	Href   string    `json:"url"`
	Baidu  []string  `json:"baidu"`
	Magnet []string  `json:"magnets"`
	Time   time.Time `json:"time"`
}

func init() {
	var err error
	client, err = mongo.NewClient(mgoURL)
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
}

// Collection -
func Collection(key string) (*mongo.Collection, error) {
	c := client.Database(mgoDB).Collection(key)
	return c, nil
}
