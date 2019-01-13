package main

import (
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const mgoURL = "mongodb://127.0.0.1:27017/"

var session *mgo.Session

// Item -
type Item struct {
	ID     bson.ObjectId   `bson:"_id,omitempty"`
	Title  string   `json:"title"`
	Href   string   `json:"url"`
	Baidu  []string `json:"baidu"`
	Magnet []string `json:"magnets"`
	Time   time.Time   `json:"time"`
}

func init() {
	var err error
	session, err = mgo.Dial(mgoURL)
	if err != nil {
		panic(err)
	}
}

func save(item *Item) error {
	c := session.DB("spider").C("taohua")
	err := c.Insert(item)
	if err != nil {
		return err
	}

	return nil
}
