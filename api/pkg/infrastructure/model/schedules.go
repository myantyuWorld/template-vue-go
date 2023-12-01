package dbmodel

import (
	"time"
)

type Schedules struct {
	PetId     uint64
	Title     string
	Date      time.Time
	Location  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (mdl *Schedules) TableName() string {
	return "schedules"
}
