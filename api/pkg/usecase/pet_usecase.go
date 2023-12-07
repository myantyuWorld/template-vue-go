package usecase

import (
	"api/pkg/domain/model"
	"api/pkg/domain/repository"
	"api/pkg/infrastructure"
	dbmodel "api/pkg/infrastructure/model"
	"fmt"
	"log"
	"strconv"
)

//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
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
	petSummary := model.PetSummary{}

	pet, err := pu.petRepository.Get(db, petId)
	if err != nil {
		return nil, err
	}
	// petDetailのコンバートに失敗した場合、エラーとする
	petDetail, err := convertPetDetail(pet)
	if err != nil {
		return nil, err
	}
	log.Println(petDetail)
	petSummary.Detail = petDetail

	// 最新のスケジュールを取得
	schedule, err := pu.scueduleRepository.FindNewest(db, petId)
	if err != nil {
		// 登録がないことを示すものをPetSummaryに設定
		petSummary.Schedules = nil
	} else {
		log.Println(fmt.Printf("%#v", schedule))
		petSchedule, _ := convertPetSchedule(schedule)
		petSummary.Schedules = petSchedule
	}

	// 最新の体調を取得
	condition, err := pu.conditionRepository.FindNewest(db, petId)
	if err != nil {
		// 登録がないことを示すものをPetSummaryに設定
		petSummary.Status = &model.PetStatus{
			Food:   nil,
			Swaet:  nil,
			Toilet: nil,
		}
	} else {
		log.Println(fmt.Printf("%#v", condition))
	}

	// TODO: dbModelのPetと、レスポンス用のpetSummaryの変換(Interface層に渡す際の処理)
	log.Println(fmt.Printf("%#v", pet))

	return &petSummary, nil
}

func convertPetSchedule(schedule *dbmodel.Schedules) (*model.ScheduledEvent, error) {
	return &model.ScheduledEvent{
		Title:     schedule.Title,
		Location:  schedule.Location,
		Timestamp: schedule.CreatedAt,
	}, nil
}

func convertPetDetail(dbmodelPets *dbmodel.Pets) (*model.Pet, error) {
	petSex, err := model.NewPetSex(strconv.FormatUint(dbmodelPets.Sex, 10))
	if err != nil {
		return nil, err
	}

	return &model.Pet{
		ID:       dbmodelPets.ID,
		Name:     dbmodelPets.Name,
		Sex:      petSex,
		Weight:   float64(dbmodelPets.NowWeight),
		Birthday: dbmodelPets.BirthDay,
	}, nil
}

func NewPetUseCase(
	pr repository.PetRepository,
	sr repository.ScheduleRepository,
	cr repository.ConditionRepository,
) PetUseCase {
	return &petUseCase{
		petRepository:       pr,
		scueduleRepository:  sr,
		conditionRepository: cr,
	}
}
