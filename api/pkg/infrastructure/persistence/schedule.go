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
	log.Println("infrastructure#persistence#schedule.go#FindNewest")

	var schedule *dbmodel.Schedules
	result := db.
		Order("id desc").
		Where("pet_id = ?", petId).
		First(&schedule)
	if result.RowsAffected == 0 {
		return nil, errors.New("検索できませんでした")
	}

	return schedule, nil
}

// Finds implements repository.ScheduleRepository.
func (*schedulePersistence) Finds(db *infrastructure.RDB, petId uint64) ([]*dbmodel.Schedules, error) {
	log.Println("infrastructure#persistence#schedule.go#Finds")

	var schedules []*dbmodel.Schedules
	result := db.
		Where("pet_id = ?", petId).
		Find(&schedules)
	if result.RowsAffected == 0 {
		return nil, errors.New("検索できませんでした")
	}

	return schedules, nil
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
