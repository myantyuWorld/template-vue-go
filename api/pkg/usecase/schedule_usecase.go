package usecase

import (
	"api/pkg/domain/model"
	"api/pkg/domain/repository"
	"api/pkg/infrastructure"
)

type ScheduleUseCase interface {
	Get(db *infrastructure.RDB, petId uint64) ([]*model.ScheduledEvent, error)
	Post(db *infrastructure.RDB, schedule *model.ScheduledEvent) (*model.ScheduledEvent, error)
}

type scheduleUseCase struct {
	scheduleRepository repository.ScheduleRepository
}

// Get implements ScheduleUseCase.
func (*scheduleUseCase) Get(db *infrastructure.RDB, petId uint64) ([]*model.ScheduledEvent, error) {
	panic("unimplemented")
}

// Post implements ScheduleUseCase.
func (*scheduleUseCase) Post(db *infrastructure.RDB, schedule *model.ScheduledEvent) (*model.ScheduledEvent, error) {
	panic("unimplemented")
}

func NewScheduleUseCase(repository repository.ScheduleRepository) ScheduleUseCase {
	return &scheduleUseCase{
		scheduleRepository: repository,
	}
}
