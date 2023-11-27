package repository

import (
	"api/pkg/domain/model"
	"api/pkg/infrastructure"
)

type PetRepository interface {
	Get(db *infrastructure.RDB, petId uint64) (*model.PetSummary, error)
}
