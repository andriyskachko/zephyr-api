package repositories

import (
	"time"

	"github.com/bzick/tokenizer"
)

type TokenType string

var (
	Word        TokenType = "word"
	Number      TokenType = "number"
	Punctuation TokenType = "punctuation"
)

var TokenTypeMap = map[tokenizer.TokenKey]TokenType{
	tokenizer.TokenKeyword:        Word,
	tokenizer.TokenInteger:        Number,
	tokenizer.TokenFloat:          Number,
	tokenizer.TokenString:         Word,
	tokenizer.TokenStringFragment: Word,
	tokenizer.TokenUnknown:        Punctuation,
}

type Token struct {
	ID        string    `db:"id"`
	Content   string    `db:"content"`
	Type      TokenType `db:"type"`
	CreatedAt time.Time `db:"created_at"`
}

func NewToken(content string, tokenType TokenType) Token {
	return Token{
		Content: content,
		Type:    tokenType,
	}
}
