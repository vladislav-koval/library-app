package book

import (
	"context"
	"errors"
	"fmt"
	"homework/helpers"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	BooksNameAuthorUniqueConstraint = "books_name_author_unique"
)

type Repository interface {
	Create(ctx context.Context, book *Book) error
	Update(ctx context.Context, book *Book) error
	Delete(ctx context.Context, id int) error

	GetByID(ctx context.Context, id int) (*Book, error)
	List(ctx context.Context, filter BookFilter) ([]Book, error)
}

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db}
}

func (r *PostgresRepository) Create(ctx context.Context, book *Book) error {
	sqlQuery := `
		INSERT INTO books 
		    (name, author, total_pages, read)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at;
	`

	err := r.db.QueryRow(
		ctx,
		sqlQuery,
		book.Name,
		book.Author,
		book.TotalPages,
		book.Read,
	).Scan(
		&book.ID,
		&book.CreatedAt,
	)

	if helpers.IsUniqueViolation(err, BooksNameAuthorUniqueConstraint) {
		return ErrBookAlreadyExists
	}

	return err
}

func (r *PostgresRepository) Update(ctx context.Context, book *Book) error {
	sqlQuery := `
		UPDATE books
		SET name = $1, author = $2, total_pages = $3, read = $4, read_at = $5
		WHERE id = $6;
	`

	cmdTag, err := r.db.Exec(
		ctx,
		sqlQuery,
		book.Name,
		book.Author,
		book.TotalPages,
		book.Read,
		book.ReadAt,
		book.ID,
	)

	if helpers.IsUniqueViolation(err, BooksNameAuthorUniqueConstraint) {
		return ErrBookAlreadyExists
	}

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return ErrBookNotFound
	}

	return nil
}

func (r *PostgresRepository) Delete(ctx context.Context, id int) error {
	sqlQuery := `
		DELETE FROM books
		WHERE id = $1;
	`

	cmdTag, err := r.db.Exec(ctx, sqlQuery, id)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return ErrBookNotFound
	}

	return nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, id int) (*Book, error) {
	sqlQuery := `
		SELECT id, name, author, total_pages, read, created_at, read_at
		FROM books
		WHERE id = $1;
	`

	row := r.db.QueryRow(ctx, sqlQuery, id)

	b := &Book{}

	if err := row.Scan(&b.ID, &b.Name, &b.Author, &b.TotalPages, &b.Read, &b.CreatedAt, &b.ReadAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrBookNotFound
		}
		return nil, err
	}

	return b, nil
}

func (r *PostgresRepository) List(ctx context.Context, filters BookFilter) ([]Book, error) {
	conditions := make([]string, 0)
	args := make([]any, 0)

	if filters.Name != "" {
		args = append(args, filters.Name)
		conditions = append(conditions, fmt.Sprintf("name = $%d", len(args)))
	}

	if filters.Author != "" {
		args = append(args, filters.Author)
		conditions = append(conditions, fmt.Sprintf("author = $%d", len(args)))
	}

	if filters.Read != nil {
		args = append(args, *filters.Read)
		conditions = append(conditions, fmt.Sprintf("read = $%d", len(args)))
	}

	sqlQuery := `
   		SELECT id, name, author, total_pages, read, created_at, read_at
   		FROM books
	`

	if len(conditions) > 0 {
		sqlQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	sqlQuery += " ORDER BY created_at DESC"

	rows, err := r.db.Query(ctx, sqlQuery, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	books := make([]Book, 0)

	for rows.Next() {
		var book Book

		err = rows.Scan(
			&book.ID,
			&book.Name,
			&book.Author,
			&book.TotalPages,
			&book.Read,
			&book.CreatedAt,
			&book.ReadAt,
		)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
