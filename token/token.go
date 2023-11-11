package token

import (
	"strings"
	"unicode"
)

type ContentType int

const (
    Word ContentType = iota
    Punctuation
    Space
    NewLine
    Unknown
)

type Token struct {
    ID          string `db:"id"`
    Content     string `db:"content"`
    ContentType ContentType `db:"content_type"`
}

func NewToken(content string) *Token {
    contentType := parseTokenContentType(content)
    return &Token{
        Content: content,
        ContentType: contentType,
    }
}

func parseTokenContentType(token string) ContentType {
    if isWord(token) {
        return Word
    } else if isPunctuation(token) {
        return Punctuation
    } else if isSpace(token) {
        return Space
    } else if isNewLine(token) {
        return NewLine
    }

    return Unknown
}

func isWord(token string) bool {
	for _, char := range token {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

func isPunctuation(token string) bool {
	punctuationMarks := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	return len(token) == 1 && strings.ContainsAny(token, punctuationMarks)
}

func isSpace(token string) bool {
	return token == " " || token == "\t"
}

func isNewLine(token string) bool {
	return token == "\n" || token == "\r\n"
}

