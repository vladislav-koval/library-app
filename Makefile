include .env
export

DB_URL=postgres://${PG_USER}:${PG_PASSWORD}@${PG_HOST}:${PG_PORT}/${PG_DB}?sslmode=${PG_SSLMODE}

run:
	go run main.go

migrate-make:
	migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	migrate -path migrations -database ${DB_URL} up

migrate-down:
	migrate -path migrations -database ${DB_URL} down
