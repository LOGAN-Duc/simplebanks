package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	db "simplebanks/db/sqlc"
	"simplebanks/util"
	"testing"
)

var testStore db.Store
var testDb *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	testDb, err = sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("could not connect to db", err)
	}
	testStore = db.NewStore(testDb)
	os.Exit(m.Run())
}
