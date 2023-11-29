package repository

import (
	"api/pkg/infrastructure"
	dbmodel "api/pkg/infrastructure/model"
)

type PetRepository interface {
	Get(db *infrastructure.RDB, petId uint64) (*dbmodel.Pets, error)
}
