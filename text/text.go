package text

import "github.com/andriyskachko/zephyr-api/token"

type Text struct {
    ID      string `db:"id"`
    Title   string `db:"title"`
    Content []token.Token `db:"content"`
}

