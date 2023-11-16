migrateup:
	migrate -path db/migrations -database "postgresql://admin:123@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://admin:123@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrateup1:
	migrate -path db/migrations -database "postgresql://admin:123@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migrations -database "postgresql://admin:123@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -count=1 -v -cover -race ./...

server:
	go run cmd/main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go go-webapp/db/sqlc Store

dockercompose:
	docker compose up --force-recreate --detach --build go-webapp

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

evans:
	evans --host localhost --port 8001 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: migrateup migratedown sqlc test server mock migrateup1 migratedown1 dockercompose proto redis evans