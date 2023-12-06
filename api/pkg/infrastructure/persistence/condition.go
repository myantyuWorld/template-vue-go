package persistence

import (
	"api/pkg/domain/repository"
	"api/pkg/infrastructure"
	dbmodel "api/pkg/infrastructure/model"
	"errors"
)

type conditionPersistence struct{}

// Post implements repository.ConditionRepository.
func (*conditionPersistence) Post(db *infrastructure.RDB, petId uint64, condition *dbmodel.Conditions) (*dbmodel.Conditions, error) {
	condition.PetId = petId
	result := db.Create(condition)
	if result.Error != nil {
		return nil, errors.New("体調履歴を追加できませんでした")
	}
	return condition, nil
}

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
