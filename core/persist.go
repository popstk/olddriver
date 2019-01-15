package core

import (
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

const (
	mgoDB  = "spider"
	mgoURL = "mongodb://127.0.0.1:27017/"
)

// Item -
type Item struct {
	Title  string    `json:"title"`
	Href   string    `json:"url"`
	Baidu  []string  `json:"baidu"`
	Magnet []string  `json:"magnets"`
	Time   time.Time `json:"time"`
}

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

