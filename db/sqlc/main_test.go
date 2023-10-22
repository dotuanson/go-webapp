package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"go-webapp/util"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBEngine, config.DBSource)
	testQueries = New(conn)

	os.Exit(m.Run())
}
