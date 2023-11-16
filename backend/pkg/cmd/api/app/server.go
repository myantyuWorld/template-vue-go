package app

import (
	"github.com/ayumi-sk2720/animal_healthcare_app/backend/pkg/interface/api"
	"github.com/ayumi-sk2720/animal_healthcare_app/backend/pkg/interface/api/handler"
)

func NewAPIServer(interactors *interactors) *api.APIServer {
	server := api.NewAPIServer()
	routing(server, interactors)

	return server
}

func routing(server *api.APIServer, interactors *interactors) {
	specifiedPetAPI := server.Group("/pet/:petID")
	specifiedPetAPI.GET("", handler.MakeGetPetSummaryHandler(interactors.GetPetSummary))
}
