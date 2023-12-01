package persistence

import (
	"api/pkg/domain/repository"
	"api/pkg/infrastructure"
	dbmodel "api/pkg/infrastructure/model"
)

type conditionPersistence struct{}

// FindNewest implements repository.ConditionRepository.
func (*conditionPersistence) FindNewest(db *infrastructure.RDB, petId uint64) (*dbmodel.Conditions, error) {
	panic("unimplemented")
}

// Finds implements repository.ConditionRepository.
func (*conditionPersistence) Finds(db *infrastructure.RDB, petId uint64) ([]*dbmodel.Conditions, error) {
	panic("unimplemented")
}

func NewConditionPersistence() repository.ConditionRepository {
	return &conditionPersistence{}
}
