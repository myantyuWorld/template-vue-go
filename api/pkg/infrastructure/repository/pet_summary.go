/***
* データベースとの入出力をする層
***/
package repository

import (
	"api/pkg/domain/model"
	"context"
)

type petSummary struct {
	*RDB
}

func NewPetSummary(db *RDB) *petSummary {
	return &petSummary{db}
}

func (repo *petSummary) Get(ctx context.Context, petId uint64) (*model.PetSummary, error) {
	// TODO: 実装
	return nil, nil
}
