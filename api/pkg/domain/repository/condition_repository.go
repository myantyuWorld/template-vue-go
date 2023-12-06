package repository

import (
	"api/pkg/infrastructure"
	dbmodel "api/pkg/infrastructure/model"
)

//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
type ConditionRepository interface {
	Post(db *infrastructure.RDB, petId uint64, condition *dbmodel.Conditions) (*dbmodel.Conditions, error)
	Finds(db *infrastructure.RDB, petId uint64) ([]*dbmodel.Conditions, error)
	FindNewest(db *infrastructure.RDB, petId uint64) (*dbmodel.Conditions, error)
}
