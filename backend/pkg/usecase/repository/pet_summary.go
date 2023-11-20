package repository

import (
	"context"

	"github.com/ayumi-sk2720/animal_healthcare_app/backend/pkg/domain/model"
)

type PetSummary interface {
	Get(ctx context.Context, petID uint64) (*model.PetSummary, error)
}
