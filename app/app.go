package app

import (
	"context"
	"fmt"

	"github.com/andriyskachko/zephyr-api/repositories"
)

func RunRepositoryDemo(ctx context.Context, repository repositories.TextRepository) {
	fmt.Println("1. MIGRATE REPOSITORY")
	// if err := textRepository.Migrate(ctx); err != nil {
	// 	log.Fatal(err)
	// }
}
