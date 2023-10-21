package main

import (
	_ "github.com/lib/pq"
	"go-webapp/db"
	"go-webapp/internal/user/repository/pgsql"
	"go-webapp/internal/user/service"
	"go-webapp/internal/user/transport/api"
	"go-webapp/router"
	"log"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize database connection: %s", err)
	}
	userRep := pgsql.NewPgSQLRepository(dbConn.GetDB())
	userSvc := service.NewService(userRep)
	userHandler := api.NewHandler(userSvc)

	//hub := ws.NewHub()
	//wsHandler := ws.NewHandler(hub)
	//go hub.Run()

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8080")
}
