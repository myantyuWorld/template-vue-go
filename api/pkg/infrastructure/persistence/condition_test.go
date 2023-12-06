package persistence_test

import (
	dbmodel "api/pkg/infrastructure/model"
	"api/pkg/infrastructure/persistence"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConditionsPost(t *testing.T) {
	db, err := OpenDb()
	if err != nil {
		log.Fatal(err)
	}
	_time := time.Now()
	_mockCondition := dbmodel.Conditions{
		Food:      uint64(1),
		Sweat:     uint64(2),
		Toilet:    uint64(3),
		CreatedAt: _time,
		UpdatedAt: _time,
	}
	t.Run("異常系_PetId_外部制約エラー", func(t *testing.T) {
		_, err := persistence.NewConditionPersistence().Post(db, uint64(99), &_mockCondition)
		require.Equal(t, errors.New("体調履歴を追加できませんでした"), err)
	})
	t.Run("正常系", func(t *testing.T) {
		output, err := persistence.NewConditionPersistence().Post(db, uint64(1), &_mockCondition)
		require.NoError(t, err)
		require.Equal(t, &_mockCondition, output)
	})
}
