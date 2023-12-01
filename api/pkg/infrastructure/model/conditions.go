package dbmodel

import "time"

type Conditions struct {
	ID        uint64
	PetId     uint64
	Food      uint64
	Sweat     uint64
	Toilet    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (mdl *Conditions) TableName() string {
	return "conditions"
}
