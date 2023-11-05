package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	db "go-webapp/db/sqlc"
	"go-webapp/internal/payment/transport/api"
	"go-webapp/util"
	"log"
)

func main() {
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.AppAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
