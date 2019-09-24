package core

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mgoURL   = "mongodb://localhost:27017"
	redisURL = "redis://localhost:6379"
)

const (
	DBName = "spider"
)

type CTime time.Time

func (c CTime) Time() time.Time {
	return time.Time(c)
}

func (c *CTime) Set(t time.Time) {
	*c = CTime(t)
}

func (c CTime) MarshalJSON() (data []byte, err error) {
	str := time.Time(c).Format("2006-01-02 15:04:05")
	return []byte(fmt.Sprintf(`"%s"`, str)), nil
}

func (c *CTime) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	t, err := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	if err != nil {
		return err
	}

	c.Set(t)
	return nil
}

// Config -
type Config struct {
	Last     CTime  `json:"last"`
	Forum    string `json:"forum"`
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
}

func (c Config) MarshalBinary() (data []byte, err error) {
	return json.MarshalIndent(c, "", " ")
}

func (c *Config) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

type Persist struct {
	key   string
	mgoDB string
	mgo   *mongo.Client
	redis *redis.Client
}

func (p *Persist) Insert(href string, item *Item) error {
	collection := p.mgo.Database(p.mgoDB).Collection(p.key)

	opt := options.FindOneAndUpdate()
	opt.SetUpsert(true)

	return collection.FindOneAndUpdate(nil, bson.M{
		"href": href,
	}, bson.M{
		"$set": item}, opt).Err()
}

func (p *Persist) Keys() ([]string, error) {
	return p.redis.Keys(fmt.Sprintf("crawler:%s*", p.key)).Result()
}

func (p *Persist) Conf(key string) (*Config, error) {
	data, err := p.redis.Get(key).Result()
	if err == redis.Nil {
		return &Config{}, nil
	}

	if err != nil {
		return nil, err
	}

	var conf Config
	if err = conf.UnmarshalBinary([]byte(data)); err != nil {
		return nil, fmt.Errorf("Persist.Conf: %w", err)
	}

	return &conf, nil
}

func (p *Persist) SaveConf(key string, conf *Config) error {
	return p.redis.Set(key, conf, 0).Err()
}

func NewPersist(key string) (*Persist, error) {
	mgo, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mgoURL))
	if err != nil {
		return nil, err
	}

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	return &Persist{
		key:   key,
		mgo:   mgo,
		redis: redis.NewClient(opt),
		mgoDB: DBName,
	}, nil
}
