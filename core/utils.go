package core

import (
	"time"
)

// TimeRange -
type TimeRange struct {
	Min    time.Time
	Max    time.Time
	Layout string
}

// NewTimeRange -
func NewTimeRange(layout string) (*TimeRange, error) {
	return &TimeRange{
		Min:    time.Unix(1<<63-1, 0),
		Max:    time.Unix(0, 0),
		Layout: layout,
	}, nil
}

// AddTime -
func (f *TimeRange) AddTime(value string) error {
	t, err := time.Parse(f.Layout, value)
	if err != nil {
		return err
	}
	f.Add(t)
	return nil
}

// Add -
func (f *TimeRange) Add(t time.Time) {
	if t.Before(f.Min) {
		f.Min = t
	}

	if t.After(f.Max) {
		f.Max = t
	}
}

// Today return today
func Today() time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}