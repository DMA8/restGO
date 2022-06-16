package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"testTask/internal/entrypoint"
	"testTask/internal/repository"
	"testTask/internal/usecases"

	"testTask/pkg/psql"
)

const (
	envDBURL            = "CONNSTR"
	envUPMigrateSQLPath = "UPMIGRATE2"
	envPort             = "PORT"
)

func main() {
	log.Println("STARTING")
	conn, err := psql.ConnectToDB(os.Getenv(envDBURL))
	if err != nil {
		fmt.Println(os.Getenv(envDBURL))
		log.Fatal("db connection bad ", err)
	}
	defer conn.Close()
	// err = psql.MakeMigrationsUp(os.Getenv(envDBURL), os.Getenv(envUPMigrateSQLPath))
	// err = psql.MakeMigrationsDown(os.Getenv(envDBURL), os.Getenv(envUPMigrateSQLPath))

	// fmt.Println(os.Getenv(envUPMigrateSQLPath))
	// err = psql.MigrateLibUp(os.Getenv(envDBURL), os.Getenv(envUPMigrateSQLPath))
	// if err != nil {
	// 	log.Fatal("migration bad ", err)
	// }
	repo := repository.NewRepository(conn)
	useCase := usecases.NewUseCase(repo)
	handler := entrypoint.NewHandler(useCase)
	router := entrypoint.NewRouter(handler)
	if err := http.ListenAndServe(os.Getenv(envPort), router); err != nil {
		log.Fatal(err)
	}
}
