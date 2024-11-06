postgres:
	docker run --name postgres17 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres:latest

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres17 dropdb simple_bank

createmigrations:
	migrate create -ext sql -dir db/migration -seq dbname

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb createmigrations sqlc