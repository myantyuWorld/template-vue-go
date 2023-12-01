package repository

import (
	"api/pkg/infrastructure"
	dbmodel "api/pkg/infrastructure/model"
)

type ConditionRepository interface {
	Finds(db *infrastructure.RDB, petId uint64) ([]*dbmodel.Conditions, error)
	FindNewest(db *infrastructure.RDB, petId uint64) (*dbmodel.Conditions, error)
}
