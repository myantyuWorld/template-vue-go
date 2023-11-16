package app

import (
	"github.com/ayumi-sk2720/animal_healthcare_app/backend/pkg/application"
	"github.com/ayumi-sk2720/animal_healthcare_app/backend/pkg/infrastructure/repository"
	iRepository "github.com/ayumi-sk2720/animal_healthcare_app/backend/pkg/usecase/repository"
)

type container struct {
	PetSummaryRepository iRepository.PetSummary
	TimestampFunc        application.TimestampFunc
}

func Inject(db *repository.RDB) *container {
	return &container{
		PetSummaryRepository: repository.NewPetSummary(db),
		TimestampFunc:        application.NewCurrentTimestampFunc(),
	}
}
