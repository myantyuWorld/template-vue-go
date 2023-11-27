package usecase

import (
	"api/pkg/domain/model"
	"context"
)

type PetUsecase interface {
	Get(ctx context.Context, petId uint64) (*model.PetSummary, error)
}
