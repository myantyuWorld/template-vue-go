package usecase_test

import (
	mock_usecase "api/mock/usecase"
	"api/pkg/domain/model"
	"api/pkg/infrastructure"
	"api/pkg/infrastructure/persistence"
	"api/pkg/usecase"
	"fmt"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestPetUsecaseGetPetSummary(t *testing.T) {
	db, err := infrastructure.ConnRDB()
	if err != nil {
		log.Fatal(err)
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUseCase := mock_usecase.NewMockPetUseCase(ctrl)
	mockPetSummary := model.PetSummary{
		Detail: &model.Pet{
			ID: uint64(1),
		},
	}
	pr := persistence.NewPetPersistence()
	sr := persistence.NewSchedulePersistence()
	cr := persistence.NewConditionPersistence()
	usecase := usecase.NewPetUseCase(
		pr,
		sr,
		cr,
	)
	t.Run("正常系_スケジュール、体調の最新が全て取得できた", func(t *testing.T) {
		mockUseCase.EXPECT().GetPetSummary(db, uint64(1)).Return(&mockPetSummary, nil)
		mockUseCase.GetPetSummary(db, uint64(1))

		fmt.Printf("%#v\n", mockPetSummary.Detail)
		t.Log("result::", mockPetSummary)
		t.Log("err::", err)

		output, err := usecase.GetPetSummary(db, uint64(1))

		fmt.Printf("%#v\n", output)
		fmt.Printf("%#v\n", err)
	})
	t.Run("正常系_スケジュールのみ、最新が取得できた", func(t *testing.T) {})
	t.Run("正常系_体調履歴のみ、最新が取得できた", func(t *testing.T) {})
	t.Run("異常系_ペットが検索できなかったケース", func(t *testing.T) {})
}
