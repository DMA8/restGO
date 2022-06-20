package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"testTask/internal/entrypoint"
	"testTask/internal/repository"
	"testTask/internal/usecases"

	"testTask/pkg/psql"
)

const (
	envDBURL      = "CONNSTR"
	envMigrations = "MIGRATES"
	envPort       = "PORT"
)

var migrateDown = flag.Bool("migrateDown", false, "executes migrateDown and exit app")

func main() {
	log.Println("app is starting")
	flag.Parse() // looking for -migrateDown flag
	log.Println("migrations begin")
	errMigrate := psql.Migrate(*migrateDown)
	if *migrateDown {
		if errMigrate != nil {
			log.Fatal(errMigrate)
		}
		log.Println("MigrateDown succeed")
		os.Exit(1)
	} else if errMigrate != nil && errMigrate.Error() != "no change" {
		log.Fatal(errMigrate)
	}
	log.Println("connection to PSQL is establishing")
	conn, err := psql.ConnectToDB(os.Getenv(envDBURL))
	if err != nil {
		log.Fatal("db connection is bad:", err)
	}
	defer conn.Close()
	
	repo := repository.NewRepository(conn)
	useCase := usecases.NewUseCase(repo)
	handler := entrypoint.NewHandler(useCase)
	router := entrypoint.NewRouter(handler)

	log.Printf("listening at %s\n", os.Getenv(envPort))
	if err := http.ListenAndServe(os.Getenv(envPort), router); err != nil {
		log.Fatal(err)
	}
}
