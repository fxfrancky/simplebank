package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/techschool/simplebank/util"
)

var testQueries *Queries

var testDB *sql.DB

func TestMain(m *testing.M) {
	// var err error
	conf, err := util.LoadConfig("../..") // "../.." the path is the root folder
	if err != nil {
		log.Fatal("Cannot log configurations")
	}
	testDB, err = sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
