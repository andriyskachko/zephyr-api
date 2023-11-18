package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgconn"
)

type PostgreSQLTextRepository struct {
	db *sql.DB
}

func NewPostgreSQLTextRepository(db *sql.DB) *PostgreSQLTextRepository {
	return &PostgreSQLTextRepository{
		db: db,
	}
}

func (r *PostgreSQLTextRepository) Migrate(ctx context.Context) error {
	query := `
    CREATE TABLE texts(
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        title VARCHAR(255),
        created_at TIMESTAMP,
    );

    CREATE TABLE tokens(
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        content VARCHAR(255),
        type VARCHAR(255),
        created_at TIMESTAMP,
    );

    CREATE TABLE text_tokens(
        text_id VARCHAR(255),
        token_id VARCHAR(255),
        PRIMARY KEY (text_id, token_id),
        FOREIGN KEY (text_id) REFERENCES texts(id) ON DELETE CASCADE,
        FOREIGN KEY (token_id) REFERENCES tokens(id) ON DELETE CASCADE
    );
    `
	_, err := r.db.ExecContext(ctx, query)
	return err
}

func (r *PostgreSQLTextRepository) Create(ctx context.Context, text Text) (*Text, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	var id string
	err = tx.QueryRowContext(ctx, "INSERT INTO texts(title) VALUES($1) RETURNING id", text.Title).Scan(&id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *PostgreSQLTextRepository) All(ctx context.Context) ([]Text, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM texts")
	all := []Text{}
	if err != nil {
		return all, err
	}
	defer rows.Close()

	for rows.Next() {
		var text Text
		if err := rows.Scan(&text.ID, &text.Title, &text.Content); err != nil {
			return all, err
		}
		all = append(all, text)
	}

	return all, nil
}

func (r *PostgreSQLTextRepository) GetByTitle(ctx context.Context, title string) (*Text, error) {
	row := r.db.QueryRowContext(ctx, "SELECT * FROM texts WHERE title = $1", title)

	var text Text
	if err := row.Scan(&text.ID, &text.Title, &text.Content); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExist
		}

		return nil, err
	}

	return &text, nil
}

func (r *PostgreSQLTextRepository) GetById(ctx context.Context, id string) (*Text, error) {
	row := r.db.QueryRowContext(ctx, "SELECT * FROM texts WHERE id = $1", id)

	var text Text
	if err := row.Scan(&text.ID, &text.Title, &text.Content); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExist
		}

		return nil, err
	}

	return &text, nil
}

func (r *PostgreSQLTextRepository) Update(ctx context.Context, id string, updated Text) (*Text, error) {
	res, err := r.db.ExecContext(ctx, "UPDATE tokens SET title = $1, content = $2 WHERE id = $3", updated.Title, updated.Content, id)
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

func (r *PostgreSQLTextRepository) Delete(ctx context.Context, id string) error {
	res, err := r.db.ExecContext(ctx, "DELETE FROM texts WHERE id = $1", id)
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

func (r *PostgreSQLTextRepository) GetTokenByContent(ctx context.Context, content string) (*Token, error) {
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
