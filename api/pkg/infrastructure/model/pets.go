package dbmodel

import "time"

type Pets struct {
	ID        uint64
	Name      string
	Birthday  time.Time
	CreatedAt time.Time
	UpdateAt  time.Time
}

func (mdl *Pets) TableName() string {
	return "pets"
}
