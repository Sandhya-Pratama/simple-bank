<<<<<<< HEAD
migrateup:
migrate -path db/migration -database "postgres://root:mysecretpassword@localhost:5433/simple_bank?sslmode=disable" -verbose up

go instal :
- migrate go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
- sqlc go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

docker :

- start db: docker start postgres17
- end db: docker end postgres17
- memeriksa status container: docker ps
- menghapus : docker rm postgres17
=======
Hello
>>>>>>> 7d787680d15e5a1552c92fcced82b4f926592b7e
