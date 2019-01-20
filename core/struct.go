package core

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	pb "github.com/popstk/olddriver/backend"
)

// Item -
type Item struct {
	Title  string    `json:"title"`
	Href   string    `json:"url"`
	Baidu  []string  `json:"baidu"`
	Magnet []string  `json:"magnets"`
	Time   time.Time `json:"time"`
}

// ToProtoItem -
func ToProtoItem(item *Item) (*pb.Item, error) {
	ts, err := ptypes.TimestampProto(item.Time)
	if err != nil {
		return nil, err
	}

	return &pb.Item{
		Title:  item.Title,
		Href:   item.Href,
		Baidu:  item.Baidu,
		Magnet: item.Magnet,
		Time:   ts,
	}, nil
}

// FromProtoItem -
func FromProtoItem(item *pb.Item) (*Item, error) {
	ts, err := ptypes.Timestamp(item.Time)
	if err != nil {
		return nil, err
	}

	return &Item{
		Title:  item.Title,
		Href:   item.Href,
		Baidu:  item.Baidu,
		Magnet: item.Magnet,
		Time:   ts,
	}, nil
}
