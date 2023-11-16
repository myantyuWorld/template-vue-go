package application

import "time"

type TimestampFunc func() time.Time

func NewCurrentTimestampFunc() TimestampFunc {
	return func() time.Time {
		return time.Now()
	}
}

func NewSpecifiedTimestampFunc(timestamp time.Time) TimestampFunc {
	return func() time.Time {
		return timestamp
	}
}
