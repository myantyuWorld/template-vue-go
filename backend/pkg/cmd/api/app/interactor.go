package app

import "github.com/ayumi-sk2720/animal_healthcare_app/backend/pkg/usecase/interactor"

type interactors struct {
	GetPetSummary interactor.GetPetSummary
}

func NewInteractors(container *container) *interactors {
	return &interactors{
		GetPetSummary: interactor.NewGetPetSummaryInteractor(container.PetSummaryRepository),
	}
}
