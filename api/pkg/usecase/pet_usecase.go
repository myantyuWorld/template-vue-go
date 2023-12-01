package usecase

import (
	"api/pkg/domain/model"
	"api/pkg/domain/repository"
	"api/pkg/infrastructure"
	"fmt"
	"log"
)

type PetUseCase interface {
	GetPetSummary(db *infrastructure.RDB, petId uint64) (*model.PetSummary, error)
}

type petUseCase struct {
	petRepository       repository.PetRepository
	scueduleRepository  repository.ScheduleRepository
	conditionRepository repository.ConditionRepository
}

// Get implements PetUseCase.
func (pu *petUseCase) GetPetSummary(db *infrastructure.RDB, petId uint64) (*model.PetSummary, error) {
	log.Println(petId)
	pet, err := pu.petRepository.Get(db, petId)
	if err != nil {
		// petテーブルが検索できない場合は、以降の処理をしても意味がないため、ここで早期リターン
		return nil, err
	}
	// 最新のスケジュールを取得
	schedule, err := pu.scueduleRepository.FIndNewest(db, petId)
	if err != nil {
		// 登録がないことを示すものをPetSummaryに設定
	} else {
		log.Println(fmt.Printf("%#v", schedule))
	}

	// 最新の体調を取得
	condition, err := pu.conditionRepository.FindNewest(db, petId)
	if err != nil {
		// 登録がないことを示すものをPetSummaryに設定
	} else {
		log.Println(fmt.Printf("%#v", condition))
	}

	// TODO: dbModelのPetと、レスポンス用のpetSummaryの変換(Interface層に渡す際の処理)
	log.Println(fmt.Printf("%#v", pet))

	return nil, nil
}

func NewPetUseCase(pr repository.PetRepository) PetUseCase {
	return &petUseCase{
		petRepository: pr,
	}
}
