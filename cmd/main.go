package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/andriyskachko/zephyr-api/app"
	"github.com/andriyskachko/zephyr-api/repositories"
	"github.com/gorilla/mux"
)

var CONTEXT_TIMEOUT = 10 * time.Second

func main() {
	loadEnv()

	db := connectToPg()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), CONTEXT_TIMEOUT)
	defer cancel()

	textRepository := repositories.NewPostgreSQLTextRepository(db)
	app.RunRepositoryDemo(ctx, textRepository)
	textsController := initTextsController(ctx, textRepository)
	r := mux.NewRouter()
	r.HandleFunc("/texts", textsController.TextGETHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
