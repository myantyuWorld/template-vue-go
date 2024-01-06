package repository

import (
	"api/pkg/infrastructure"
	dbmodel "api/pkg/infrastructure/model"
)

//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
type PetRepository interface {
	Get(db *infrastructure.RDB, petId uint64) (*dbmodel.Pets, error)
}
