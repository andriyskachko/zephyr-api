package app

import (
	"context"
	"fmt"
	"log"

	"github.com/andriyskachko/zephyr-api/text"
	"github.com/andriyskachko/zephyr-api/token"
)

func RunRepositoryDemo(ctx context.Context, textRepository text.TextRepository, tokenRepository token.TokenRepository) {
    fmt.Println("1. MIGRATE TEXT REPOSITORY")
    if err := textRepository.Migrate(ctx); err != nil {
        log.Fatal(err)
    }

    fmt.Println("2. MIGRATE TOKEN REPOSITORY")
    if err := tokenRepository.Migrate(ctx); err != nil {
        log.Fatal(err)
    }
}

