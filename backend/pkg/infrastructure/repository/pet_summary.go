package repository

import (
	"context"

	"github.com/ayumi-sk2720/animal_healthcare_app/backend/pkg/domain/model"
)

type petSummary struct {
	*gormRepsotiroy
}

func NewPetSummary(db *RDB) *petSummary {
	return &petSummary{newGormRepository(db)}
}

func (repo *petSummary) Get(ctx context.Context, petID uint64) (*model.PetSummary, error) {
	return nil, nil
}
