package test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/Sandhya-Pratama/simple-bank/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:mysecretpassword@localhost:5433/simple_bank?sslmode=disable"
)

var testQueries *sqlc.Queries

func TestMain(m *testing.M) {
	// Membuat koneksi pool menggunakan pgxpool
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	defer conn.Close() // Menutup koneksi pool setelah testing selesai

	testQueries = sqlc.New(conn)

	os.Exit(m.Run())
}
