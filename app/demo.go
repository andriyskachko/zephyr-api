package app

import (
	"context"
	"fmt"
	"log"

	"github.com/andriyskachko/zephyr-api/text"
)

func RunRepositoryDemo(ctx context.Context, textRepository text.TextRepository) {
    fmt.Println("1. MIGRATE REPOSITORY")
    if err := textRepository.Migrate(ctx); err != nil {
        log.Fatal(err)
    }

}

