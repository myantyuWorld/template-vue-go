package persistence

import (
	"api/pkg/domain/repository"
	"api/pkg/infrastructure"
	dbmodel "api/pkg/infrastructure/model"
	"errors"
	"log"
)

type schedulePersistence struct{}

// FIndNewest implements repository.ScheduleRepository.
func (*schedulePersistence) FIndNewest(db *infrastructure.RDB, petId uint64) (*dbmodel.Schedules, error) {
	log.Println("infrastructure#persistence#schedule.go#")
	panic("unimplemented")
}

// Finds implements repository.ScheduleRepository.
func (*schedulePersistence) Finds(db *infrastructure.RDB, petId uint64) ([]*dbmodel.Schedules, error) {
	log.Println("infrastructure#persistence#schedule.go#")
	panic("unimplemented")
}

// Post implements repository.ScheduleRepository.
func (*schedulePersistence) Post(db *infrastructure.RDB, petId uint64, schedule *dbmodel.Schedules) (*dbmodel.Schedules, error) {
	log.Println("infrastructure#persistence#schedule.go#")

	schedule.PetId = petId
	result := db.Create(schedule)
	if result.Error != nil {
		return nil, errors.New("スケジュールを追加できませんでした")
	}
	return schedule, nil
}

func NewSchedulePersistence() repository.ScheduleRepository {
	return &schedulePersistence{}
}
