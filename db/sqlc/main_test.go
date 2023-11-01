package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go-webapp/util"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}
