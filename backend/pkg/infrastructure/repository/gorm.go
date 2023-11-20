package repository

import (
	"gorm.io/gorm"
)

type gormRepsotiroy struct {
	*gorm.DB
}

func newGormRepository(db *RDB) *gormRepsotiroy {
	return &gormRepsotiroy{db.DB}
}
