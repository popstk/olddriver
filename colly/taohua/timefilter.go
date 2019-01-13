package main
import (
	"time"
)

// TimeRange -
type TimeRange struct {
	Min time.Time
	Max time.Time
	Layout string
}

// NewTimeRange -
func NewTimeRange(layout, now string) (*TimeRange, error){
	t := time.Now()
	if now != "" {
		var err error
		t, err = time.Parse(layout, now)
		if err != nil {
			return nil, err
		}
	}

	return &TimeRange{
		Min: t,
		Max: t,
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
func (f *TimeRange) Add(t  time.Time) {
	if t.Before(f.Min) {
		f.Min = t
	}

	if t.After(f.Max) {
		f.Max = t
	}
}

// BeforeMin - 
func (f *TimeRange) BeforeMin(t time.Time) bool {
	return f.Min.Before(t)
}

