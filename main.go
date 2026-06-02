package main

import (
	"context"
	"fmt"
	"homework/book"
	"homework/book_http"
	"homework/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, cfg config.PostgresConfig) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, cfg.ConnString())
}

func main() {

	//  TODO: Добавить тест на constraintю Лучше integration test на repo или хотя бы unit test для helpers.IsUniqueViolation.

	ctx := context.Background()
	cfg := config.Load()

	pool, err := NewPool(ctx, cfg.Postgres)

	if err != nil {
		panic(err)
	}
	defer pool.Close()

	repo := book.NewRepo(pool)
	service := book.NewService(repo)
	handlers := book_http.NewHandlers(service)
	server := book_http.NewServer(handlers)

	if err := server.StartServer(); err != nil {
		fmt.Println("error while start server", err)
	}
}
