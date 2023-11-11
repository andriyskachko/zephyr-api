package text

import (
	"reflect"
	"regexp"
	"strings"
	"time"
)

type Text struct {
    ID      string `db:"id"`
    Title   string `db:"title"`
    Content []Token `db:"content"`
    CreatedAt time.Time `db:"created_at"`
}

func NewText(title string, text string) *Text {
    content := tokenize(text)

    return &Text{
        Title: title,
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

func tokenize(text string) []Token {
	var tokens []Token

	wordRegex := regexp.MustCompile(`\b\w+\b`)
	punctuationRegex := regexp.MustCompile(`[.,;!?]`)
	letterRegex := regexp.MustCompile(`\b\w\b`)

	start := 0
	for {
		// Find the next match for word, punctuation, or letter
		wordMatch := wordRegex.FindStringIndex(text[start:])
		punctuationMatch := punctuationRegex.FindStringIndex(text[start:])
		letterMatch := letterRegex.FindStringIndex(text[start:])

		// Find the minimum of the three matches
		minMatch := minNumsInArray(wordMatch, punctuationMatch, letterMatch)
		if minMatch == nil {
			break
		}

		// Adjust the index to the global text index
		minMatch[0] += start
		minMatch[1] += start

		// Extract the token content and type based on the match
		content := text[minMatch[0]:minMatch[1]]
		var tokenType TokenType
        if reflect.DeepEqual(minMatch, wordMatch) {
            tokenType = Word
        } else if reflect.DeepEqual(minMatch, punctuationMatch) {
            tokenType = Punctuation
        } else if reflect.DeepEqual(minMatch, letterMatch) {
            tokenType = Letter
        }

		// Append the token to the result
		tokens = append(tokens, Token{Content: content, Type: tokenType})

		// Move the start index to the end of the match
		start = minMatch[1]
	}

	return tokens
}

func minNumsInArray(nums ...[]int) []int {
	if len(nums) == 0 {
		return nil
	}

	min := nums[0]
	for _, num := range nums[1:] {
		if num[0] < min[0] {
			min = num
		}
	}

	return min
}
