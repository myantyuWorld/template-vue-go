package handler

import (
	"api/pkg/infrastructure"
	"api/pkg/usecase"
	"log"

	"github.com/labstack/echo/v4"
)

type PetHandler interface {
	HandlePetGet() echo.HandlerFunc
}

type petHandler struct {
	petUsecase usecase.PetUseCase
	db         infrastructure.RDB
}

// HandlePetGet implements PetHandler.
func (ph petHandler) HandlePetGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("aaa")
		var pet interface{}
		var err error

		if pet, err = ph.petUsecase.Get(&ph.db, 1); err != nil {
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
