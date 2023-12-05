package repository

import (
	"api/pkg/infrastructure"
	dbmodel "api/pkg/infrastructure/model"
)

//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
type ScheduleRepository interface {
	Finds(db *infrastructure.RDB, petId uint64) ([]*dbmodel.Schedules, error)
	FIndNewest(db *infrastructure.RDB, petId uint64) (*dbmodel.Schedules, error)
	Post(db *infrastructure.RDB, petId uint64, schedule *dbmodel.Schedules) (*dbmodel.Schedules, error)
}
