package db

import (
	"database/sql"
	"log"
	"os"
	"simplebank/util"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("can't load config: ", err)
	}
	dbDriver := config.DBDriver
	dbSource := config.DBSource
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can connect to database: ", err)
	}

	testQueries = New(testDB)

	// Start running unitest
	os.Exit(m.Run())
}
