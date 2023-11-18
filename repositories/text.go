package repositories

import (
	"strings"
	"time"

	"github.com/bzick/tokenizer"
)

type Text struct {
	ID        string                       `db:"id" json:"id"`
	Title     string                       `db:"title" json:"title"`
	Content   map[tokenizer.TokenKey]Token `db:"content" json:"content"`
	CreatedAt time.Time                    `db:"created_at" json:"created_at"`
}

func NewText(title string, text string) *Text {
	content := tokenizeText(text)

	return &Text{
		Title:   title,
		Content: content,
	}
}

func (t *Text) ToString() string {
	var builder strings.Builder

	totalLength := 0
	for _, txt := range t.Content {
		totalLength += len(txt.Content)
	}
	builder.Grow(totalLength)

	for _, txt := range t.Content {
		builder.WriteString(txt.Content)
	}

	return builder.String()
}

func tokenizeText(text string) map[tokenizer.TokenKey]Token {
	parser := tokenizer.New()
	parser.SetWhiteSpaces([]byte{})
	stream := parser.ParseString(text)
	defer stream.Close()
	tokens := make(map[tokenizer.TokenKey]Token)

	for stream.IsValid() {
		currentToken := stream.CurrentToken()
		tokenContent := currentToken.ValueString()
		tokenType := TokenTypeMap[currentToken.Key()]
		token := NewToken(tokenContent, tokenType)
		tokens[tokenizer.TokenKey(currentToken.ID())] = token
		stream.GoNext()
	}

	return tokens
}
