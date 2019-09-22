package core

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	pb "github.com/popstk/olddriver/backend"
)

// Item -
type Item struct {
	Title  string    `json:"title"`
	URL    string    `json:"url"`
	Tag    string    `json:"tag"`
	Time   time.Time `json:"time"`
	Baidu  []string  `json:"baidu" bson:"baidu,omitempty"`
	Magnet []string  `json:"magnet" bson:"magnet,omitempty"`
	Link   []string  `json:"link" bson:"link,omitempty"`
}

// ToProtoItem -
func ToProtoItem(item *Item) (*pb.Item, error) {
	ts, err := ptypes.TimestampProto(item.Time)
	if err != nil {
		return nil, err
	}

	return &pb.Item{
		Url:    item.URL,
		Tag:    item.Tag,
		Title:  item.Title,
		Time:   ts,
		Baidu:  item.Baidu,
		Magnet: item.Magnet,
		Link:   item.Link,
	}, nil
}

// FromProtoItem -
func FromProtoItem(item *pb.Item) (*Item, error) {
	ts, err := ptypes.Timestamp(item.Time)
	if err != nil {
		return nil, err
	}

	return &Item{
		URL:    item.Url,
		Tag:    item.Tag,
		Title:  item.Title,
		Time:   ts,
		Baidu:  item.Baidu,
		Magnet: item.Magnet,
		Link:   item.Link,
	}, nil
}
