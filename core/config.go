package core

import (
	"os"
	"time"
	"os/signal"
	"syscall"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

const (
	mgoConfig  = "config"
	configFile = "config.json"
)

// Config -
type Config struct {
	MongoURL string
}

// SpiderConfig -
type SpiderConfig struct {
	Name string
	Cron string
	Last time.Time
}

// GetSpiderConfig -
func GetSpiderConfig(key string) (*SpiderConfig, error) {
	var conf SpiderConfig

	c, err := Collection(mgoConfig)
	if err != nil {
		return nil, err
	}

	err = c.FindOne(nil, bson.M{"name": key}).Decode(&conf)
	if err == mongo.ErrNoDocuments {
		return &SpiderConfig{
			Name: key,
			Cron: "0 0 22 * * *",
			Last: Today(),
		}, nil
	}

	if err != nil {
		return nil, err
	}

	return &conf, nil
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
		"name": key,
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