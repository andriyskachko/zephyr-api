package token

type ContentType int

const (
    Word ContentType = iota
    Punctuation
    Space
    NewLine
)

type Token struct {
    ID          string `db:"id"`
    Content     string `db:"content"`
    ContentType ContentType `db:"content_type"`
}

