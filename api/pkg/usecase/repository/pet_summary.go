package repository

import (
	"api/pkg/domain/model"
	"context"
)

type PetSummary interface {
	Get(ctx context.Context, petId uint64) (*model.PetSummary, error)
}
