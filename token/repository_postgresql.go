package token

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgconn"
)

type PostgreSQLTokenRepository struct {
    db *sql.DB
}

func NewPostgreSQLTokenRepository(db *sql.DB) *PostgreSQLTokenRepository {
    return &PostgreSQLTokenRepository{
        db: db,
    }
}

func (r *PostgreSQLTokenRepository) Migrate(ctx context.Context) error {
    query := `
    CREATE TABLE tokens(
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        content VARCHAR(255) NOT NULL,
        type INTEGER NOT NULL
    );
    `
    _, err := r.db.ExecContext(ctx, query)
    return err
}

func (r *PostgreSQLTokenRepository) Create(ctx context.Context, token Token) (*Token, error) {
    var id string
    err := r.db.QueryRowContext(ctx, "INSERT INTO tokens(content, type) VALUES($1, $2) RETURNING id", token.Content, token.Type).Scan(&id); if err != nil {
        var pgxError *pgconn.PgError
        if errors.As(err, &pgxError) {
            if pgxError.Code == RecordExistsErrorCode {
                return nil, ErrDuplicate
            }
        }
        return nil, err
    }
    token.ID = id

    return &token, nil
}

func (r *PostgreSQLTokenRepository) All(ctx context.Context) ([]Token, error) {
    rows, err := r.db.QueryContext(ctx, "SELECT * FROM tokens")
    all := []Token{}
    if err != nil {
        return all, err
    }
    defer rows.Close()

    for rows.Next() {
        var token Token
        if err := rows.Scan(&token.ID, &token.Content, &token.Type); err != nil {
            return all, err
        }
        all = append(all, token)
    }

    return all, nil
}

func (r *PostgreSQLTokenRepository) GetByContent(ctx context.Context, content string) (*Token, error) {
    row := r.db.QueryRowContext(ctx, "SELECT * FROM tokens WHERE content = $1", content)

    var token Token
    if err := row.Scan(&token.ID, &token.Content, &token.Type); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrNotExist
        }

        return nil, err
    }

    return &token, nil
}

func (r *PostgreSQLTokenRepository) GetById(ctx context.Context, id string) (*Token, error) {
    row := r.db.QueryRowContext(ctx, "SELECT * FROM tokens WHERE id = $1", id)

    var token Token
    if err := row.Scan(&token.ID, &token.Content, &token.Type); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrNotExist
        }

        return nil, err
    }

    return &token, nil
}

func (r *PostgreSQLTokenRepository) Update(ctx context.Context, id string, updated Token) (*Token, error) {
    res, err := r.db.ExecContext(ctx, "UPDATE tokens SET content = $1, type = $2 WHERE id = $3", updated.Content, updated.Type, id)
    if err != nil {
        var pgxError *pgconn.PgError
        if errors.As(err, &pgxError) {
            if pgxError.Code == RecordExistsErrorCode {
                return nil, ErrDuplicate
            }
        }
        return nil, err
    }

    rowsAffected, err := res.RowsAffected()
    if err != nil {
        return nil, err
    }

    if rowsAffected == 0 {
        return nil, ErrUpdateFailed
    }

    return &updated, nil
}

func (r *PostgreSQLTokenRepository) Delete(ctx context.Context, id string) error {
    res, err := r.db.ExecContext(ctx, "DELETE FROM tokens WHERE id = $1", id)
    if err != nil {
        return err
    }

    rowsAffected, err := res.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        return ErrDeleteFailed
    }

    return err
}

