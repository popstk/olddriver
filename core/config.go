package core

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

const (
	mgoConfig  = "config"
	configFile = "config.json"
)

// ErrNoDocuments -
var ErrNoDocuments = mongo.ErrNoDocuments

// Config -
type Config struct {
	MongoURL string
}

// SpiderConfig -
type SpiderConfig struct {
	ID         primitive.ObjectID `bson:"_id"`
	Collection string             `bson:"collection"`
	Tag        string             `bson:"tag"`
	Cron       string             `bson:"cron"`
	Last       time.Time          `bson:"last"`
}

// NewSpiderConfig -
func NewSpiderConfig(c, tag string) *SpiderConfig {
	return &SpiderConfig{
		Collection: c,
		Tag:        tag,
		Cron:       "0 0 22 * * *",
		Last:       Today(),
	}
}

// GetSpiderConfig -
func GetSpiderConfig(name string) ([]*SpiderConfig, error) {
	c, err := Collection(mgoConfig)
	if err != nil {
		return nil, err
	}

	cursor, err := c.Find(nil, bson.D{{"collection", name}}, nil)
	if err != nil {
		return nil, err
	}

	confs := make([]*SpiderConfig, 0)
	for cursor.Next(context.Background()) {
		var conf SpiderConfig
		err = cursor.Decode(&conf)
		if err != nil {
			return nil, err
		}

		confs = append(confs, &conf)
	}

	return confs, nil
}

// SaveSpiderConfig -
func SaveSpiderConfig(key string, conf *SpiderConfig) error {
	c, err := Collection(mgoConfig)
	if err != nil {
		return err
	}

	opt := options.FindOneAndUpdate()
	opt.SetUpsert(true)

	c.FindOneAndUpdate(nil, bson.M{
		"_id": conf.ID,
	}, bson.M{
		"$set": conf,
	}, opt)

	return nil
}

// WaitForExit -
func WaitForExit() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	<-ch
}
