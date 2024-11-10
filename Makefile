postgres:
	docker run --name postgres17 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres:latest

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres17 dropdb simple_bank

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

.PHONY: postgres createdb dropdb createmigrations migrateup migratedown sqlc test server mock

