package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ayumi-sk2720/animal_healthcare_app/backend/pkg/domain/model"
	"github.com/ayumi-sk2720/animal_healthcare_app/backend/pkg/interface/api"
	"github.com/ayumi-sk2720/animal_healthcare_app/backend/pkg/usecase/interactor"
	"github.com/labstack/echo/v4"
)

func MakeGetPetSummaryHandler(usecase interactor.GetPetSummary) echo.HandlerFunc {
	return func(c echo.Context) error {
		in, err := newGetPetSummaryInput(c)
		if err != nil {
			return err
		}

		return usecase.Get(api.NewContext(c), in, newGetPetSummaryPresenter(c, time.Now()))
	}
}

func newGetPetSummaryInput(c echo.Context) (interactor.GetPetSummaryInput, error) {
	id, err := strconv.ParseUint(c.Param("petID"), 10, 64)
	if err != nil {
		return nil, err
	}

	return &getPetSummaryInput{
		petID: id,
	}, nil
}

type getPetSummaryInput struct {
	petID uint64
}

func (in *getPetSummaryInput) PetID() uint64 {
	return in.petID
}

func newGetPetSummaryPresenter(c echo.Context, timestamp time.Time) interactor.GetPetSummaryPresenter {
	return &getPetSummaryPresenter{
		c:         c,
		timestamp: timestamp,
	}
}

type getPetSummaryPresenter struct {
	c         echo.Context
	timestamp time.Time
}

func (presenter *getPetSummaryPresenter) Render(summary *model.PetSummary) error {
	return presenter.c.JSON(http.StatusOK, newGetPetSummaryResponse(summary))
}

func newGetPetSummaryResponse(summary *model.PetSummary) *getPetSummaryResponse {
	return &getPetSummaryResponse{}
}

type getPetSummaryResponse struct {
	BaseInfo struct {
		Name         string `json:"name"`
		Age          string `json:"age"`
		Sex          string `json:"sex"`
		NowWeight    string `json:"now_weight"`
		TargetWeight string `json:"target_weight"`
		Birthday     string `json:"birthday"`
	} `json:"baseInfo"`
	DosageSchedules struct {
		Today struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
			Date  string `json:"date"`
		} `json:"today"`
		Next struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
			Date  string `json:"date"`
		} `json:"next"`
	} `json:"dosage_schedules"`
	PhysicalCondition struct {
		Food   int `json:"food"`
		Sweat  int `json:"sweat"`
		Toilet int `json:"toilet"`
	} `json:"physical_condition"`
	Memo struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		Date  string `json:"date"`
	} `json:"memo"`
	Schedules []struct {
		ID       int    `json:"id"`
		Date     string `json:"date"`
		Title    string `json:"title"`
		Location string `json:"location"`
	} `json:"schedules"`
}
