package persistence_test

import (
	dbmodel "api/pkg/infrastructure/model"
	"api/pkg/infrastructure/persistence"
	mock_repository "api/pkg/mock/repository"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestPost(t *testing.T) {
	db, err := OpenDb()
	if err != nil {
		log.Fatal(err)
	}
	t.Run("異常系", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		_time := time.Now()
		_mockSchedule := dbmodel.Schedules{
			PetId:     uint64(1),
			Title:     "ホゲホゲ", // 現状、Mysqlの設定が至らなく、日本語をinsertしようとした時にエラーとなる、それを利用してみる
			Date:      _time,
			Location:  "fugafuga",
			CreatedAt: _time,
			UpdatedAt: _time,
		}
		output, err := persistence.NewSchedulePersistence().Post(db, uint64(1), &_mockSchedule)
		log.Print(output)
		require.Equal(t, errors.New("スケジュールを追加できませんでした"), err)
	})
	t.Run("正常系", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockScheduleRepository(ctrl)
		_time := time.Now()
		_mockSchedule := dbmodel.Schedules{
			PetId:     uint64(1),
			Title:     "hogehoge",
			Date:      _time,
			Location:  "fugafuga",
			CreatedAt: _time,
			UpdatedAt: _time,
		}
		// モックを設定
		mockRepo.EXPECT().Post(db, uint64(1), &_mockSchedule).Return(&_mockSchedule, nil)
		// 実際にモックを呼び出す | https://qiita.com/gold-kou/items/81562f9142323b364a60
		mockRepo.Post(db, uint64(1), &_mockSchedule)

		output, _ := persistence.NewSchedulePersistence().Post(db, uint64(1), &_mockSchedule)
		require.Equal(t, &_mockSchedule, output)
	})
}

func TestFindNewest(t *testing.T) {
	db, err := OpenDb()
	if err != nil {
		log.Fatal(err)
	}
	t.Run("異常系", func(t *testing.T) {
		output, err := persistence.NewSchedulePersistence().FindNewest(db, uint64(99))
		log.Print(output)
		require.Equal(t, errors.New("検索できませんでした"), err)
	})
	t.Run("正常系", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		_mockSchedule := dbmodel.Schedules{
			PetId: uint64(1),
			Title: "hogehoge",
		}

		output, _ := persistence.NewSchedulePersistence().FindNewest(db, uint64(1))
		require.Equal(t, _mockSchedule.Title, output.Title)
	})
}

func TestFinds(t *testing.T) {
	db, err := OpenDb()
	if err != nil {
		log.Fatal(err)
	}
	t.Run("異常系", func(t *testing.T) {
		output, err := persistence.NewSchedulePersistence().Finds(db, uint64(99))
		log.Print(output)
		require.Equal(t, errors.New("検索できませんでした"), err)
	})
	t.Run("正常系", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockScheduleRepository(ctrl)
		var _mockSchedule []*dbmodel.Schedules
		mockRepo.EXPECT().Finds(db, uint64(1)).Return(&_mockSchedule, nil)
		mockRepo.Finds(db, uint64(1))

		output, _ := persistence.NewSchedulePersistence().Finds(db, uint64(1))
		require.Equal(t, len(_mockSchedule), len(output))
	})
}
