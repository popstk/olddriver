package core

import (
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// TimeRange -
type TimeRange struct {
	Min    time.Time
	Max    time.Time
	Layout string
}

// NewTimeRange -
// real max time
// https://stackoverflow.com/questions/25065055/what-is-the-maximum-time-time-in-go/32620397
func NewTimeRange(layout string) (*TimeRange, error) {
	return &TimeRange{
		Min:    time.Unix(1<<63-62135596801, 999999999),
		Max:    time.Unix(0, 0),
		Layout: layout,
	}, nil
}

// AddTime -
func (f *TimeRange) AddTime(value string) (time.Time, error) {
	t, err := time.ParseInLocation(f.Layout, value, time.Local)
	if err != nil {
		return time.Now(), err
	}
	f.Add(t)
	return t, nil
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

// WaitForExit -
func WaitForExit() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	<-ch
}

// JoinFullURL
func JoinFullURL(main, path string) string {
	u, _ := url.Parse(path)
	base, _ := url.Parse(main)
	return base.ResolveReference(u).String()
}

func JoinURL(main *url.URL, path string) string {
	u, _ := url.Parse(path)
	return main.ResolveReference(u).String()
}
