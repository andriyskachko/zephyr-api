package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/andriyskachko/zephyr-api/app"
	"github.com/andriyskachko/zephyr-api/text"
	"github.com/andriyskachko/zephyr-api/token"
	"github.com/joho/godotenv"

    _ "github.com/jackc/pgx/v4/stdlib"
)

var contextTimeout = 10 * time.Second

func main() {
    loadEnv()

    db := connectToPg()
    defer db.Close()

    textRepository := text.NewPostgreSQLTextRepository(db)
    tokenRepository := token.NewPostgreSQLTokenRepository(db)

    ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
    defer cancel()

    app.RunRepositoryDemo(ctx, textRepository, tokenRepository)
}

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
        pgDb = os.Getenv("POSTGRES_DB")
    )

    db, err := sql.Open("pgx", fmt.Sprintf("postgres://postgres:%s@localhost:%s/%s", pgPassword, dockerPort, pgDb))
    if err != nil {
        log.Fatal(err)
    }

    return db
}
