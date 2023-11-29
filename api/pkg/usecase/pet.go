package usecase

import (
	"api/pkg/domain/model"
	"api/pkg/domain/repository"
	"api/pkg/infrastructure"
	"fmt"
	"log"
)

type PetUseCase interface {
	Get(db *infrastructure.RDB, petId uint64) (*model.PetSummary, error)
}

type petUseCase struct {
	petRepository repository.PetRepository
}

// Get implements PetUseCase.
func (pu *petUseCase) Get(db *infrastructure.RDB, petId uint64) (*model.PetSummary, error) {
	log.Println(petId)
	pet, err := pu.petRepository.Get(db, petId)
	if err != nil {
		return nil, err
	}
	// dbModelのPetと、レスポンス用のpetSummaryの変換(Interface層に渡す際の処理)
	log.Println(fmt.Printf("%#v", pet))

	return nil, nil
}

func NewPetUseCase(pr repository.PetRepository) PetUseCase {
	return &petUseCase{
		petRepository: pr,
	}
}
