package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/andriyskachko/zephyr-api/controllers"
	"github.com/andriyskachko/zephyr-api/repositories"
	"github.com/andriyskachko/zephyr-api/services"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func connectToPg() *sql.DB {
	var (
		pgPassword = os.Getenv("POSTGRES_PASSWORD")
		dockerPort = os.Getenv("DOCKER_PORT_MAPPING")
		pgDb       = os.Getenv("POSTGRES_DB")
	)

	db, err := sql.Open("pgx", fmt.Sprintf("postgres://postgres:%s@localhost:%s/%s", pgPassword, dockerPort, pgDb))
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping the database")
	}

	return db
}

func initTextsController(ctx context.Context, textRepository repositories.TextRepository) *controllers.TextController {
	textsService := services.NewTextService(ctx, textRepository)
	textsController := controllers.NewTextController(*textsService)

	return textsController
}
