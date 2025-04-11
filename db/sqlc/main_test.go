package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"simplebank/db/utils"
	"testing"
)

//const (
//	dbDriver = "postgres"
//	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
//)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, er := utils.LoadConfig("../..")
	if er != nil {
		log.Fatal("Loading Config Failed", er)
	}

	var err error
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("COULD NOT CONNECT TO THE DATABASE: ", err)
	}
	defer testDB.Close()

	testQueries = New(testDB)
	os.Exit(m.Run())
}
