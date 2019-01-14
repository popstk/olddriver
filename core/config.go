package core

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

const (
	mgoConfig = "config"
)

// SpiderConfig -
type SpiderConfig struct {
	Name string
	Last time.Time
}

// GetLastTime -
func GetLastTime(key string) (time.Time, error) {
	c := client.Database(mgoDB).Collection(mgoConfig)

	var conf SpiderConfig
	err := c.FindOne(nil, bson.M{"name": key}).Decode(&conf)
	if err == mongo.ErrNoDocuments {
		return time.Now(), nil
	}

	if err != nil {
		return time.Now(), err
	}

	return conf.Last, nil
}
