package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	db "go-webapp/db/sqlc"
	"go-webapp/internal/payment/transport/api"
	"go-webapp/internal/payment/transport/rpc"
	"go-webapp/pb"
	"go-webapp/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
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
	runGrpcServer(config, store)
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := rpc.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGoWebAppServer(grpcServer, server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatalf("cannot create listener")
	}
	log.Printf("gRPC server listening at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start gRPC server")
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
