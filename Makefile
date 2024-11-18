postgres:
	docker run --name postgres --network bank-network -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=sandhya123 -p 5432:5432 -d postgres:16

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

createmigrations:
	migrate create -ext sql -dir db/migration -seq dbname

migrateup:
	migrate -path db/migration -database "postgresql://root:mysecretpassword@localhost:5433/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:mysecretpassword@localhost:5433/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go  github.com/Sandhya-Pratama/simple-bank/db/sqlc Store

build : docker build -t simplebank:latest .

run : docker run --name simplebank -p 8080:8080 -e GIN_MODE=release  simplebank:latest

.PHONY: postgres createdb dropdb createmigrations migrateup migratedown sqlc test server mock build run

