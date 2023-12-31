package handler

import (
	"api/pkg/infrastructure"
	"api/pkg/usecase"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PetHandler interface {
	HandlePetGet() echo.HandlerFunc
	HandleScheduleGet() echo.HandlerFunc
	HandleSchedulePost() echo.HandlerFunc
}

type petHandler struct {
	petUsecase usecase.PetUseCase
	db         infrastructure.RDB
}

// HandleScheduleGet implements PetHandler.
func (handler *petHandler) HandleScheduleGet() echo.HandlerFunc {
	panic("unimplemented")
}

// HandleSchedulePost implements PetHandler.
func (handler *petHandler) HandleSchedulePost() echo.HandlerFunc {
	panic("unimplemented")
}

// HandlePetGet implements PetHandler.
func (ph petHandler) HandlePetGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		var pet interface{}
		var err error

		petId, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			log.Println("parse error")
			return c.JSON(400, err)
		}

		if pet, err = ph.petUsecase.GetPetSummary(&ph.db, uint64(petId)); err != nil {
			log.Println("handler error")
			log.Println(err)
			return c.JSON(500, err)
			// return c.JSON(500, response.Error(c.Response().Writer, http.StatusInternalServerError, err, "internal server error"))
		}
		// return response.JSON(c.Response().Writer, http.StatusOK, pet)
		return c.JSON(200, pet)
	}
}

func NewPetHandler(pu usecase.PetUseCase, db infrastructure.RDB) PetHandler {
	log.Print("NewPetHandler")
	return &petHandler{
		petUsecase: pu,
		db:         db,
	}
}
