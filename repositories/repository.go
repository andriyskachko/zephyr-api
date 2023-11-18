package repositories

import (
	"context"
	"errors"
)

var (
	ErrDuplicate          = errors.New("text already exists")
	ErrNotExist           = errors.New("text does not exist")
	ErrUpdateFailed       = errors.New("update text failed")
	ErrDeleteFailed       = errors.New("delete text failed")
	RecordExistsErrorCode = "23505"
)

type TextRepository interface {
	Migrate(ctx context.Context) error
	Create(ctx context.Context, text Text) (*Text, error)
	All(ctx context.Context) ([]Text, error)
	GetByTitle(ctx context.Context, title string) (*Text, error)
	GetById(ctx context.Context, id string) (*Text, error)
	Update(ctx context.Context, id string, updated Text) (*Text, error)
	Delete(ctx context.Context, id string) error

	GetTokenByContent(ctx context.Context, content string) (*Token, error)
}
