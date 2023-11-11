package token

import (
	"context"
	"errors"
)

var (
	ErrDuplicate    = errors.New("token already exists")
	ErrNotExist     = errors.New("token does not exist")
	ErrUpdateFailed = errors.New("update token failed")
	ErrDeleteFailed = errors.New("delete token failed")
    RecordExistsErrorCode  = "23505"
)

type TokenRepository interface {
	Migrate(ctx context.Context) error
	Create(ctx context.Context, token Token) (*Token, error)
	All(ctx context.Context) ([]Token, error)
	GetById(ctx context.Context, id string) (*Token, error)
	GetByContent(ctx context.Context, content string) (*Token, error)
	Update(ctx context.Context, id string, updated Token) (*Token, error)
	Delete(ctx context.Context, id string) error
}

