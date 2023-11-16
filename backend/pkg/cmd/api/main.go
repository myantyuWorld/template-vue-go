package main

import (
	"log"

	"github.com/ayumi-sk2720/animal_healthcare_app/backend/pkg/cmd/api/app"
	"github.com/ayumi-sk2720/animal_healthcare_app/backend/pkg/infrastructure/repository"
)

func init() {

}

func main() {
	dbCfg := repository.NewMySQLConfig("localhost", 3060, "test", "user", "password")
	db, err := repository.ConnRDB(dbCfg)
	if err != nil {
		log.Fatal(err)
	}

	container := app.Inject(db)
	interacetors := app.NewInteractors(container)

	server := app.NewAPIServer(interacetors)

	if err := server.Start(":3000"); err != nil {
		log.Fatal(err)
	}
}
