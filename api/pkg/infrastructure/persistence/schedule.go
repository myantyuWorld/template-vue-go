package persistence

import (
	"api/pkg/domain/repository"
	"api/pkg/infrastructure"
	dbmodel "api/pkg/infrastructure/model"
)

type schedulePersistence struct{}

// FIndNewest implements repository.ScheduleRepository.
func (*schedulePersistence) FIndNewest(db *infrastructure.RDB, petId uint64) (*dbmodel.Schedules, error) {
	panic("unimplemented")
}

// Finds implements repository.ScheduleRepository.
func (*schedulePersistence) Finds(db *infrastructure.RDB, petId uint64) ([]*dbmodel.Schedules, error) {
	panic("unimplemented")
}

// Post implements repository.ScheduleRepository.
func (*schedulePersistence) Post(db *infrastructure.RDB, schedule *dbmodel.Schedules) (*dbmodel.Schedules, error) {
	panic("unimplemented")
}

func NewSchedulePersistence() repository.ScheduleRepository {
	return &schedulePersistence{}
}
