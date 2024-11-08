package test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Sandhya-Pratama/simple-bank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:mysecretpassword@localhost:5433/simple_bank?sslmode=disable"
)

var testQueries *sqlc.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	// Connect ke db
	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = sqlc.New(testDB)

	os.Exit(m.Run())
}
