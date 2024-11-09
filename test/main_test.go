package test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Sandhya-Pratama/simple-bank/db/sqlc"
	"github.com/Sandhya-Pratama/simple-bank/util"
	_ "github.com/lib/pq"
)

var testQueries *sqlc.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	// Load environment variable
	config, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Connect ke db
	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = sqlc.New(testDB)

	os.Exit(m.Run())
}
