package usecase

import (
	"api/pkg/domain/model"
	"api/pkg/domain/repository"
	"api/pkg/infrastructure"
)

type PetUseCase interface {
	Get(db *infrastructure.RDB, petId uint64) (*model.PetSummary, error)
}

type petUseCase struct {
	petRepository repository.PetRepository
}

// Get implements PetUseCase.
func (*petUseCase) Get(db *infrastructure.RDB, petId uint64) (*model.PetSummary, error) {
	panic("unimplemented")
}

func NewPetUseCase(pr repository.PetRepository) PetUseCase {
	return &petUseCase{
		petRepository: pr,
	}
}
