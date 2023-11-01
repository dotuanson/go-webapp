migrateup:
	migrate -path db/migrations -database "postgresql://admin:123@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://admin:123@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover -race ./...

server:
	go run cmd/main.go

.PHONY: migrateup migratedown sqlc test server