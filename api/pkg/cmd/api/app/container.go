package app

import (
	"api/pkg/application"
	"api/pkg/infrastructure/repository"
	connect "api/pkg/infrastructure/repository"
	petSummaryUsecase "api/pkg/usecase/repository"
)

type container struct {
	PetSummaryRepository petSummaryUsecase.PetSummary
	TimestampFunc        application.TimestampFunc
}

func Inject(db *connect.RDB) *container {
	return &container{
		PetSummaryRepository: repository.NewPetSummary(db),
		TimestampFunc:        application.NewCurrentTimestampFunc(),
	}
}
