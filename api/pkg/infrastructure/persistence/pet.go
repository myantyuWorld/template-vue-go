package persistence

import (
	"api/pkg/domain/model"
	"api/pkg/domain/repository"
	"api/pkg/infrastructure"
)

type petPersistence struct{}

// Get implements repository.PetRepository.
func (*petPersistence) Get(db *infrastructure.RDB, petId uint64) (*model.PetSummary, error) {
	panic("unimplemented")
}

func NewPetPersistence() repository.PetRepository {
	return &petPersistence{}
}
