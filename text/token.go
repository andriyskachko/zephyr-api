package text

import (
	"time"
)

type TokenType int

const (
    Word TokenType = iota
    Punctuation
    Letter
)

type Token struct {
    ID          string `db:"id"`
    Content     string `db:"content"`
    Type TokenType `db:"type"`
    CreatedAt time.Time `db:"created_at"`
}

func NewToken(content string, tokenType TokenType) *Token {
    return &Token{
        Content: content,
        Type: tokenType,
    }
}
