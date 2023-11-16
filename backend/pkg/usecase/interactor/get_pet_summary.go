package interactor

import (
	"context"

	"github.com/ayumi-sk2720/animal_healthcare_app/backend/pkg/domain/model"
	"github.com/ayumi-sk2720/animal_healthcare_app/backend/pkg/usecase/repository"
)

type GetPetSummaryInput interface {
	PetID() uint64
}

type GetPetSummaryPresenter interface {
	Render(output *model.PetSummary) error
}

type GetPetSummary interface {
	Get(ctx context.Context, input GetPetSummaryInput, presenter GetPetSummaryPresenter) error
}

type getPetSummary struct {
	summaryRepository repository.PetSummary
}

func NewGetPetSummaryInteractor(summaryRepository repository.PetSummary) GetPetSummary {
	return &getPetSummary{
		summaryRepository: summaryRepository,
	}
}

func (interactor *getPetSummary) Get(ctx context.Context, input GetPetSummaryInput, presenter GetPetSummaryPresenter) error {
	summary, err := interactor.summaryRepository.Get(ctx, input.PetID())
	if err != nil {
		return err
	}

	return presenter.Render(summary)
}
