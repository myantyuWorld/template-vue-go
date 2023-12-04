package dbmodel

import "time"

// insert into
//
//	pets(
//	  id,
//	  name,
//	  birth_day,
//	  sex,
//	  now_weight,
//	  target_weight,
//	  created_at,
//	  updated_at
//	)
//
// values
//
//	(1, "なつ", now(), 1, 4000, 5000, now(), now());
type Pets struct {
	ID           uint64
	Name         string
	Sex          uint64
	BirthDay     time.Time
	NowWeight    uint64
	TargetWeight uint64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (mdl *Pets) TableName() string {
	return "pets"
}
