package repository

import (
	"api/pkg/infrastructure"
	dbmodel "api/pkg/infrastructure/model"
)

type ScheduleRepository interface {
	Finds(db *infrastructure.RDB, petId uint64) ([]*dbmodel.Schedules, error)
	FIndNewest(db *infrastructure.RDB, petId uint64) (*dbmodel.Schedules, error)
	Post(db *infrastructure.RDB, schedule *dbmodel.Schedules) (*dbmodel.Schedules, error)
}
